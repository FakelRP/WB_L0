package main

import (
	consumer2 "WB_L0/internal/consumer-jetstream"
	"WB_L0/internal/handler/order"
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("launched")

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	newJS, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	streamName := "EVENTS"

	stream, err := newJS.Stream(ctx, streamName)
	if err != nil {
		log.Fatal(err)
	}

	cons, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		InactiveThreshold: 10 * time.Millisecond,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Created consumer", cons.CachedInfo().Name)

	orderHandler := order.New()

	consumer := consumer2.New(cons, orderHandler)

	ch := make(chan bool)

	err = consumer.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	ch <- true

	// graceful shutdown
	//consumer.Close()

}
