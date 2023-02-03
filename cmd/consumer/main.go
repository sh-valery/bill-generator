package main

import (
	"context"
	"fmt"
	kafka "github.com/segmentio/kafka-go"
	"time"
)

const (
	topic          = "bills"
	broker1Address = "localhost:9092"
)

func main() {
	ctx := context.Background()
	conusme(ctx)
}

func conusme(ctx context.Context) {
	// initialize the reader with the broker addresses, the topic, and the partition
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker1Address},
		Topic:     topic,
		Partition: 0,
	})

	for {
		// read a single message from the topic
		m, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		// log the message key and value
		fmt.Printf("message at offset %d: %s = %s", m.Offset, string(m.Key), string(m.Value))

		// sleep for a second
		time.Sleep(time.Second)
	}
}
