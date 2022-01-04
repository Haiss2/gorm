package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017/")
	// Connect to MongoDB
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	client, err := mongo.NewClient(clientOptions)
	client.Connect(context.TODO())
	mongo.NewIndexOptionsBuilder()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("HealCheck")
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	collection = client.Database("test_db").Collection("cars")

	var indexView *mongo.IndexView

	// Create two indexes: {name: 1, email: 1} and {name: 1, age: 1}
	// For the first index, specify no options. The name will be generated as
	// "name_1_email_1" by the driver.
	// For the second index, specify the Name option to explicitly set the name
	// to "nameAge".
	models := []mongo.IndexModel{
		{
			Keys: bson.D{{"name", 1}, {"email", 1}},
		},
		{
			Keys:    bson.D{{"name", 1}, {"age", 1}},
			Options: options.Index().SetName("nameAge"),
		},
	}

	// Specify the MaxTime option to limit the amount of time the operation can
	// run on the server
	opts := options.CreateIndexes().SetMaxTime(2 * time.Second)
	names, err := indexView.CreateMany(context.TODO(), models, opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created indexes %v\n", names)
}
