package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"fmt"
	// "github.com/joho/godotenv"
)

var DB *mongo.Database

func ConnectDB() {
    clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatalf("❌ MongoDB connection error: %v", err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatalf("❌ MongoDB ping error: %v", err)
    }

    DB = client.Database(os.Getenv("DB_NAME"))
    fmt.Println("✅ Connected to MongoDB")
}

func GetCollection(collectionName string) *mongo.Collection {
    return DB.Collection(collectionName)
}
