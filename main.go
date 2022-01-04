package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ []bson.M

type Car struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Year      int
	Make      string
	ModelName string
	DriverID  int
}

var collection *mongo.Collection

func batch(num int) []interface{} {
	makes := []string{"Toyota", "Honda", "Nissan", "Ford"}
	ModelName := []string{"Tundra", "Accord", "Sentra", "F-150"}
	batch_users := []interface{}{}
	for n := 0; n < num; n++ {
		batch_users = append(batch_users, Car{
			Year:      1970 + n%52,
			Make:      makes[rand.Int31n(4)],
			ModelName: ModelName[rand.Int31n(4)],
			DriverID:  n,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}
	return batch_users
}

func insertOne(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	makes := []string{"Toyota", "Honda", "Nissan", "Ford"}
	ModelName := []string{"Tundra", "Accord", "Sentra", "F-150"}
	for n := 0; n < num; n++ {
		collection.InsertOne(context.TODO(), Car{
			Year:      1970 + n%52,
			Make:      makes[rand.Int31n(4)],
			ModelName: ModelName[rand.Int31n(4)],
			DriverID:  n,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}

}

func simpleQuery() {
	filter := bson.D{{"make", "Honda"}}
	opts := options.Find().SetSort(bson.D{{"driverid", 1}})
	cursor, _ := collection.Find(context.TODO(), filter, opts)
	var x []bson.M
	cursor.All(context.TODO(), &x)
	fmt.Println("Found", len(x))
}

func updateDoc() {
	filter := bson.D{{"driverid", 1}}

	for i := 0; i < 10000; i++ {
		update := bson.D{{"$set", bson.D{{"year", -i}}}}
		collection.UpdateOne(context.TODO(), filter, update)
	}

	// updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func updateDocs() {
	filter := bson.D{{"make", "Honda"}}
	update := bson.D{{"$set", bson.D{{"year", -1}, {"createdat", time.Now()}}}}
	_, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
}

func deleteMany() {
	filter := bson.D{{"make", "Honda"}}
	collection.DeleteMany(context.TODO(), filter)
}

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

	start := time.Now().UnixNano()
	fmt.Println("Start time: ", start)

	// var wg sync.WaitGroup

	// for i := 0; i < 100; i++ {
	// 	wg.Add(1)
	// 	go insertOne(&wg, 100)
	// }
	// wg.Wait()

	// collection.InsertMany(context.TODO(), batch(1000000))

	// simpleQuery()

	// updateDoc()

	updateDocs()

	// deleteMany()

	end := time.Now().UnixNano()
	fmt.Println("End time: ", end)
	fmt.Println("Write time:", (end-start)/1000000, "ms")

}
