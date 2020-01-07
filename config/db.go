package config

import (
	"context"
	"log"
	"newspaper-backend/constants"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client : Database variable
var Client *mongo.Client

// DbConnect : Connecting MongoDB
func DbConnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(constants.MongoURI)
	Client, _ = mongo.Connect(ctx, clientOptions)
	if Client != nil {
		log.Println("DB Connected")
	}
}
