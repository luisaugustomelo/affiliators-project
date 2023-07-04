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

type QueuePublisher struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

func NewQueuePublisher(conn *amqp.Connection, ch *amqp.Channel, q amqp.Queue) *QueuePublisher {
	return &QueuePublisher{
		conn: conn,
		ch:   ch,
		q:    q,
	}
}

func (qp *QueuePublisher) PublishMessage(message interfaces.Message, db *gorm.DB) (*models.QueueProcessing, error) {
	body, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = qp.ch.Publish(
		"",
		qp.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	mq := &models.QueueProcessing{
		UserId:  1,
		Hash:    "",
		Status:  "pending",
		Message: "waiting to process",
	}

	if err := db.Save(mq).Error; err != nil {
		log.Print(err)
		return nil, err
	}

	return mq, nil
}

func PublishToQueue(message interfaces.Message, db *gorm.DB) (*models.QueueProcessing, error) {
	amqpHost := utils.GetEnv("AMQPHOST", "amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial(amqpHost)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return nil, err
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
		return nil, err
	}

	qp := NewQueuePublisher(conn, ch, q)
	return qp.PublishMessage(message, db)
}
