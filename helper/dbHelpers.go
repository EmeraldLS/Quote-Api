package helper

import (
	"context"
	"strconv"
	"time"

	"github.com/EmeraldLS/quote-generator/db"
	"github.com/EmeraldLS/quote-generator/model"
	"github.com/bxcodec/faker/v4"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection = db.OpenCollection()
var ctx, cancel = context.WithTimeout(context.TODO(), time.Second*10)

func CreateQuote(quote model.Quote) (interface{}, error) {
	defer cancel()

	result, err := collection.InsertOne(ctx, quote)
	if err != nil {
		return "", err
	}
	return result.InsertedID, nil
}

func Populate() interface{} {
	defer cancel()
	var result *mongo.InsertOneResult
	for i := 0; i < 50; i++ {
		var quote = model.Quote{
			Quote_Text: faker.Word(),
			Author:     faker.Name(),
		}
		result, _ = collection.InsertOne(ctx, quote)
	}

	return result
}

func GetAllQuotes() []model.Quote {
	defer cancel()
	var quotes []model.Quote
	filter := bson.M{}

	cursor, _ := collection.Find(ctx, filter)
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var quote model.Quote
		cursor.Decode(&quote)
		quotes = append(quotes, quote)
	}
	return quotes
}

func GetRelatedQuote(c *gin.Context) []model.Quote {
	defer cancel()
	var quotes []model.Quote
	filter := bson.M{}
	findOptions := options.Find()
	if s := c.Query("s"); s != "" {
		filter = bson.M{
			"quote_text": bson.M{
				"$regex": primitive.Regex{
					Pattern: s,
					Options: "i",
				},
			},
		}

	}

	if sort := c.Query("sort"); sort != "" {
		if sort == "asc" {
			findOptions.SetSort(bson.M{"author": 1})
		} else if sort == "desc" {
			findOptions.SetSort(bson.M{"author": -1})
		}
	}

	page, _ := strconv.Atoi(c.Query("page"))
	var perPage int64 = 10
	findOptions.SetSkip(int64(page-1) * perPage)
	findOptions.SetLimit(perPage)

	cursor, _ := collection.Find(ctx, filter, findOptions)
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var quote model.Quote
		cursor.Decode(&quote)
		quotes = append(quotes, quote)
	}

	return quotes
}

func DeleteAllQuotes() int64 {
	filter := bson.D{{}}
	result, _ := collection.DeleteMany(ctx, filter)
	return result.DeletedCount
}

func DeleteQuote(c *gin.Context) int64 {
	filter := bson.M{}
	if s := c.Query("s"); s != "" {
		filter = bson.M{"_id": s}
	}
	result, _ := collection.DeleteOne(ctx, filter)
	return result.DeletedCount
}

func DeleteAllSimilarQuotes(c *gin.Context) int64 {
	filter := bson.M{}
	if sim := c.Query("sim"); sim != "" {
		filter = bson.M{
			"quote_text": bson.M{
				"$regex": primitive.Regex{
					Pattern: sim,
					Options: "i",
				},
			},
		}
	}
	result, _ := collection.DeleteMany(ctx, filter)
	return result.DeletedCount
}
