package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("❌ MongoDB client creation error:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("❌ MongoDB connection error:", err)
	}

	// Ping
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ MongoDB ping error:", err)
	}

	DB = client.Database(os.Getenv("MONGO_DB"))
	log.Println("✅ MongoDB Connected to database:", DB.Name())
}

func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}
