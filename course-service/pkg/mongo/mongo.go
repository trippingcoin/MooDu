package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Conn   *mongo.Database
	Client *mongo.Client
}

func ConnectMongo() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	db := &DB{
		Conn:   client.Database("minimoodle"),
		Client: client,
	}

	err = db.Client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("ping connection mongoDB Error: %w ", err)
	}

	log.Println("Connected to MongoDB")

	return db.Conn, nil
}
