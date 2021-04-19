package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func setupDB() (*mongo.Client, error) {
	uri := getMongoURI()
	fmt.Println("USING URI: ", uri)
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}
	if err = client.Connect(context.Background()); err != nil {
		return nil, fmt.Errorf("CONNECTION ERROR: %v ", err)
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	fmt.Println("Mongodb connected successfully.")
	return client, nil
}

func getMongoURI() string {
	mongoURI := os.Getenv("MONGO_DB_URI")
	if mongoURI == "" {
		return "mongodb://localhost:27017/test_db"
	}
	return mongoURI
}
