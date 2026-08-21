package order

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func Producer(order Order) {

	data, err := json.Marshal(order)
	if err != nil {
		log.Fatal("json marshall error", err)
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:29092"},
		Topic:   "orders",
	})

	defer writer.Close()

	time.Sleep(1 * time.Second)

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(order.Product),
		Value: data,
	})

	if err != nil {
		log.Fatal("cannot write to kafka", err)
	} else {
		fmt.Println("producer sent a message")
	}
}
