package workers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/luisaugustomelo/hubla-challenge/database/models"
	"github.com/luisaugustomelo/hubla-challenge/interfaces"
	"github.com/luisaugustomelo/hubla-challenge/utils"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func PublishToQueue(message interfaces.Message, db *gorm.DB) error {
	amqpHost := utils.GetEnv("AMQPHOST", "amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial(amqpHost)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer ch.Close()

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
		return err
	}

	body, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	mq := &models.QueueProcessing{}
	mq.UserId = 1
	mq.Hash = ""
	mq.Status = "pending"
	mq.Message = "waiting to process"

	if err := db.Save(&mq).Error; err != nil {
		log.Print(err)
		return err
	}

	return nil
}
