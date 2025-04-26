package db

import (
	"context"
	"log"
	"time"
	"username-check-api/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatal(err)
	}
}

func UsernameExists(username string) bool {
	collection := client.Database("userdb").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, _ := collection.CountDocuments(ctx, bson.M{"username": username})
	return count > 0
}

func FetchAllUsernames() ([]string, error) {
	var usernames []string
	collection := client.Database("userdb").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			continue
		}
		usernames = append(usernames, user.Username)
	}

	return usernames, nil
}

func InsertUsername(username string) error {
	collection := client.Database("userdb").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, bson.M{"username": username})
	return err
}
