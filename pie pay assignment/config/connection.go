package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/event"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	monitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			fmt.Printf("\n✅ MongoDB Query: %s\n%v\n", evt.CommandName, evt.Command)
		},
		Succeeded: func(_ context.Context, evt *event.CommandSucceededEvent) {
			fmt.Printf("✅ MongoDB Query Success: %s (duration: %v)\n", evt.CommandName, evt.Duration)
		},
		Failed: func(_ context.Context, evt *event.CommandFailedEvent) {
			fmt.Printf("❌ MongoDB Query Failed: %s (err: %v)\n", evt.CommandName, evt.Failure)
		},
	}
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://PiePayAssignment:1234567890@cluster0.twptwmc.mongodb.net/").SetMonitor(monitor))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("flipkart_offers_db")
	log.Println("✅ Connected to MongoDB!")
}
