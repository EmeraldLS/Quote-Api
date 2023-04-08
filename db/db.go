package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var dbname = os.Getenv("DBNAME")
var colname = os.Getenv("COLNAME")

func DBInstance() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print(err)
	}
	var uri = os.Getenv("URI")
	clientOptions := options.Client().ApplyURI(uri)
	client, _ := mongo.Connect(ctx, clientOptions)

	return client
}

func OpenCollection() *mongo.Collection {
	client := DBInstance()
	collection = client.Database(dbname).Collection(colname)
	return collection
}
