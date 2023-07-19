package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"os"
)

func main() {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect("0.0.0.0:4222")
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Drain()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("Jetstream error - %f", err)
	}

	cfg := nats.StreamConfig{
		Name:     "EVENTS",
		Subjects: []string{"events.>"},
	}

	cfg.Storage = nats.FileStorage

	_, err = js.AddStream(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("created the stream")
}
