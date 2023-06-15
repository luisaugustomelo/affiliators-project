package queue

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/luisaugustomelo/hubla-challenge/interfaces"
	"github.com/luisaugustomelo/hubla-challenge/services"
	"github.com/streadway/amqp"
)

func ConsumerToQueue() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer ch.Close()

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

	go func() {
		for d := range msgs {
			m := &interfaces.Message{}
			err := json.Unmarshal(d.Body, m)
			if err != nil {
				fmt.Printf("Error decoding JSON: %s\n", err)
				continue
			}

			fileContent, err := base64.StdEncoding.DecodeString(m.File)
			filename, _ := services.ProcessFile(string(fileContent))
			//if err != nil {
			//	return fiberError(c, fiber.StatusInternalServerError, "Failed to process file", err)
			//}

			transactions, _ := services.ReadTransactions(filename)
			//if err != nil {
			//	return fiberError(c, fiber.StatusInternalServerError, "Failed to read transactions", err)
			//}

			//return c.Status(fiber.StatusOK).JSON(fiber.Map{"filepath": "/images/single/" + filename, "transactions": transactions})
			fmt.Printf("Received a message: %s %s %v\n", m.Email, m.File, transactions)
		}
	}()
}
