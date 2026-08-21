package order

import (
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func CreateTopic() {
	conn, err := kafka.Dial("tcp", "localhost:29092")
	if err != nil {
		log.Fatal("kafka connection error", err)
	}

	controller, err := conn.Controller()
	if err != nil {
		log.Fatal("cannot get controller", err)
	}

	controllerConn, err := kafka.Dial("tcp", fmt.Sprintf("%s:%d", controller.Host, controller.Port))
	if err != nil {
		log.Fatal("connection error to controller", err)
	}

	defer controllerConn.Close()

	topic := "orders"

	err = controllerConn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})

	if err != nil {
		log.Fatal("cannot create topic", err)
	} else {
		fmt.Println("kafka topic created")
	}
}
