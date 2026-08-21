package order

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func Consumer() {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:29092"},
		Topic:   "orders",
		GroupID: "orders-group",
	})

	defer reader.Close()

	fmt.Println("kafka listening...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		var o Order
		json.Unmarshal(msg.Value, &o)
		fmt.Printf("order received %s : %d\n", o.Product, o.Quantity)
	}

}
