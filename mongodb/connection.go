// mongodb/connection.go
package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, *mongo.Collection) {
	err := godotenv.Load() // Load variables from .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := os.Getenv("MONGO_URI")
	if connectionString == "" {
		connectionString = "mongodb://localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("cve_database")
	collection := database.Collection("cves")

	fmt.Println("Connected to MongoDB!")
	return client, collection
}
