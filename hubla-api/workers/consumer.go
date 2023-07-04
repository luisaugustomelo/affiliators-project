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

const numProcessingRoutines = 5

type MessageQueue struct {
	db   *gorm.DB
	ch   *amqp.Channel
	q    amqp.Queue
	lock sync.Mutex
}

func NewMessageQueue(db *gorm.DB, ch *amqp.Channel, q amqp.Queue) *MessageQueue {
	return &MessageQueue{
		db: db,
		ch: ch,
		q:  q,
	}
}

func (mq *MessageQueue) startProcessing() {
	msgs, err := mq.ch.Consume(
		mq.q.Name,
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

	for i := 0; i < numProcessingRoutines; i++ {
		go func() {
			for d := range msgs {
				mq.processMessage(d)
			}
		}()
	}
}

func (mq *MessageQueue) processMessage(d amqp.Delivery) {
	mq.lock.Lock()
	defer mq.lock.Unlock()

	queue, err := mq.ch.QueueInspect(mq.q.Name)
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
		return
	}

	mq.processFile(m)
}

func (mq *MessageQueue) processFile(m *interfaces.Message) {
	mqItem := models.QueueProcessing{}

	if err := mq.db.Where("user_id = ? AND status = ?", m.UserId, "pending").Last(&mqItem).Error; err != nil {
		return
	}

	mqItem.Status = "error"
	mqItem.Message = "wasn't generated"

	fileContent, err := base64.StdEncoding.DecodeString(m.File)
	if err != nil {
		mqItem.Hash = "empty"
		if err := mq.db.Save(&mqItem).Error; err != nil {
			log.Print(err)
		}
		return
	}

	filename, err := services.ProcessFile(string(fileContent))
	if err != nil {
		mqItem.Hash = filename
		if err := mq.db.Save(&mqItem).Error; err != nil {
			log.Print(err)
		}
		return
	}

	sales, balances, err := services.ReadSales(filename)
	if err != nil {
		mqItem.Hash = filename
		mqItem.Message = "error to read sales file"
		if err := mq.db.Save(&mqItem).Error; err != nil {
			log.Print(err)
		}
		return
	}

	if err := mq.db.CreateInBatches(balances, len(balances)).Error; err != nil {
		log.Print(err)
		return
	}

	for index := range sales {
		sales[index].UserID = m.UserId
		sales[index].Hash = strings.TrimSuffix(filename, ".txt")
	}
	if err := mq.db.CreateInBatches(sales, len(sales)).Error; err != nil {
		log.Print(err)
		return
	}

	mqItem.Status = "success"
	mqItem.Hash = filename
	mqItem.Message = "done"

	if err := mq.db.Save(&mqItem).Error; err != nil {
		log.Print(err)
	}

	fmt.Printf("Received a message: %s %s %v\n", m.Email, m.File, sales)
}

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

	mq := NewMessageQueue(db, ch, q)
	mq.startProcessing()
}
