package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	Url      string
	Client   *mongo.Client
	Database *mongo.Database
}

func CreateClient(url string) *Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	database := client.Database("test")

	return &Client{
		Url:      url,
		Client:   client,
		Database: database,
	}
}

func (c Client) Connect() string {
	fmt.Println("Pinging...")
	err := c.Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(3)
	}

	fmt.Println("Connected to MongoDB!")
	return "Connected to MongoDB!"
}
