package queue

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func InitQueue(queueName string) {
	go func() {

		conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			log.Fatalf("Failed to open a channel: %v", err)
		}
		defer ch.Close()

		q, err := ch.QueueDeclare(
			queueName, // Tên queue
			false,     // Durable
			true,      // Auto-Delete (xóa queue khi không còn consumer hoặc chương trình kết thúc)
			true,      // Exclusive (chỉ sử dụng queue này từ một kết nối duy nhất)
			false,     // No-Wait
			nil,       // Arguments
		)
		if err != nil {
			log.Fatalf("Failed to declare a queue: %v", err)
		}

		msgs, err := ch.Consume(
			q.Name,
			"",    // consumer
			true,  // auto-ack
			false, // exclusive
			false, // no-local
			false, // no-wait
			nil,   // args
		)

		if err != nil {
			log.Println(err)
			return
		}

		for m := range msgs {
			log.Println(string(m.Body))
		}
	}()
}
