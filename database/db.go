package database


import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

//const connectionString = "mongodb://localhost:27017"

const DBNAME = "stark"

const COLLECTION = "users"

var Collection *mongo.Collection

func init() {
	MONGODB := os.Getenv("MONGODB")
	// set client options
	clientOption := options.Client().ApplyURI(MONGODB)

	// connect to db
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("connected to mongo successfully")
	Collection = client.Database(DBNAME).Collection(COLLECTION)
	fmt.Println("connection instance created")
}

