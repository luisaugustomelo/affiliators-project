package workers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/luisaugustomelo/hubla-challenge/database/models"
	"github.com/luisaugustomelo/hubla-challenge/interfaces"
	"github.com/luisaugustomelo/hubla-challenge/services"
	"github.com/luisaugustomelo/hubla-challenge/utils"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func ConsumerToQueue(db *gorm.DB) {
	amqpHost := utils.GetEnv("AMQPHOST", "amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial(amqpHost)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}

	channelName := utils.GetEnv("CNAME", "hubla-sales-queue")
	q, err := ch.QueueDeclare(
		channelName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	var lock sync.Mutex
	go func(msgs <-chan amqp.Delivery) {
		for d := range msgs {
			lock.Lock()
			queue, err := ch.QueueInspect(q.Name)
			if err != nil {
				// handle error
				fmt.Println(err)
			} else {
				fmt.Printf("Number of pending messages in the queue: %d\n", queue.Messages)
			}

			m := &interfaces.Message{}
			err = json.Unmarshal(d.Body, m)
			if err != nil {
				fmt.Printf("Error decoding JSON: %s\n", err)
				continue
			}

			mq := models.QueueProcessing{}

			// Find currently user
			if err := db.Where("user_id = ? AND status = ?", m.UserId, "pending").Last(&mq).Error; err != nil {
				continue
			}

			mq.Status = "error"
			mq.Message = "wasn't generated"

			fileContent, err := base64.StdEncoding.DecodeString(m.File)
			if err != nil {
				mq.Hash = "empty"
				if err := db.Save(&mq).Error; err != nil {
					log.Print(err)
				}
				continue
			}
			filename, err := services.ProcessFile(string(fileContent))
			if err != nil {
				mq.Hash = filename
				if err := db.Save(&mq).Error; err != nil {
					log.Print(err)
				}
				continue
			}

			sales, balances, err := services.ReadSales(filename)
			if err != nil {
				mq.Hash = filename
				mq.Message = "error to read sales file"
				if err := db.Save(&mq).Error; err != nil {
					log.Print(err)
				}
				continue
			}

			if err := db.CreateInBatches(balances, len(balances)).Error; err != nil {
				log.Print(err)
				continue
			}

			for index := range sales {
				sales[index].UserID = m.UserId
				sales[index].Hash = strings.TrimSuffix(filename, ".txt")
			}
			if err := db.CreateInBatches(sales, len(sales)).Error; err != nil {
				log.Print(err)
				continue
			}

			mq.Status = "success"
			mq.Hash = filename
			mq.Message = "done"

			if err := db.Save(&mq).Error; err != nil {
				log.Print(err)
			}

			fmt.Printf("Received a message: %s %s %v\n", m.Email, m.File, sales)
			lock.Unlock()
		}
	}(msgs)
}
