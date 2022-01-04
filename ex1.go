package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017/")
	client, _ := mongo.NewClient(clientOptions)
	client.Connect(context.TODO())

	coll := client.Database("test_db").Collection("cars")

	x := coll.Indexes()
	fmt.Println(x.List(context.TODO()))

}
