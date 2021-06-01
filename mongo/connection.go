package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect() *mongo.Client {
	// get .env
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURL := os.Getenv("MONGO_URL")

	// log.Printf(mongoURL)

	// connect to database
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprint(mongoURL)))
	if err != nil {
		log.Fatal(err)
	}

	// if connection takes more than 10 seconds timeout and throw an error
	ctx, ctxErr := context.WithTimeout(context.Background(), 10*time.Second)
	if ctxErr != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
	// Ping the database to see if it worked
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	// Show all available databases
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	return client
}
