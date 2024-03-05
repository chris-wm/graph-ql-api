package dispatcher

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/electivetechnology/utility-library-go/logger"
	"github.com/electivetechnology/utility-library-go/pubsub"
)

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("events/dispatcher")
}

func DispatchEvent(event Event) (Event, error) {
	// Get topic prefix
	prefix := os.Getenv("GCP_PUBSUB_TOPIC_PREFIX")
	if prefix == "" {
		prefix = "d-graph-ql-api"
	}

	// Construct topic name from event
	topic := fmt.Sprintf("%s-%s", prefix, event.Name)

	// Forward to Dispatch()
	return Dispatch(event, topic)
}

func Dispatch(event Event, topic string) (Event, error) {
	// Publish asyncroniously
	var wg sync.WaitGroup
	log.Printf("Sending message to Pub/Sub topic %s: %v", topic, event.Name)
	wg.Add(1)

	// Setup Pub/Sub client
	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to Initialize Pub/Sub Client: %v", err)
	}

	go func(event Event) error {
		err = pubsubClient.Publish(topic, event.Payload)
		if err != nil {
			log.Fatalf("Failed to publish message to Pub/Sub: %v", err)
			return err
		}
		wg.Done()

		return nil
	}(event)

	// Wait for message to be publish
	wg.Wait()

	return event, nil
}
