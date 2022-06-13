package main

import (
	"context"
	"log"
	"time"

	"github.com/you/aws/internal/email"
	"github.com/you/aws/internal/pkg/cloud/aws"
)

func main() {
	// Create a cancellable context.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a session instance.
	ses, err := aws.New(aws.Config{
		Address: "http://localhost:4566",
		Region:  "us-east-1",
		Profile: "test",
		ID:      "test",
		Secret:  "test",
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Set queue URL.
	url := "http://localhost:4566/000000000000/queue1"

	// Instantiate client.
	client := aws.NewSQS(ses, time.Second*5)

	// Instantiate consumer and start consuming.
	email.NewConsumer(client, email.ConsumerConfig{
		Type:      email.AsyncConsumer,
		QueueURL:  url,
		MaxWorker: 2,
		MaxMsg:    10,
	}).Start(ctx)
}
