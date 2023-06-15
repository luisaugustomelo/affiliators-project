package workers

import (
	"encoding/json"
	"fmt"

	"github.com/luisaugustomelo/hubla-challenge/interfaces"
	"github.com/streadway/amqp"
)

func PublishToQueue(message interfaces.Message) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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

	q, err := ch.QueueDeclare(
		"hubla-sales-queue",
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

	return nil
}
