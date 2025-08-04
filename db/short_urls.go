package db

import (
	"context"
	"log"
	"strings"

	"kiku/main/lib"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ShortUrl struct {
	ID          string `bson:"shorten_id"`
	Destination string `bson:"shorten_url_destination"`
}

func GetShortenedUrls() []ShortUrl {
	collection := lib.Db("shortened_urls")

	ctx := context.TODO()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error finding shortened URLs:", err)
		return nil
	}
	defer cursor.Close(ctx)

	var results []ShortUrl
	if err := cursor.All(ctx, &results); err != nil {
		log.Println("Error decoding results:", err)
		return nil
	}

	return results
}

func ShortenUrl(url string) ShortUrl {
	collection := lib.Db("shortened_urls")
	ctx := context.TODO()

	str, err := lib.GenerateRandomString(10)
	if err != nil {
		log.Println("Error generating random string:", err)
		return ShortUrl{}
	}

	var existing ShortUrl
	err = collection.FindOne(ctx, bson.M{"shorten_id": str}).Decode(&existing)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Println("Error checking for existing ID:", err)
		return ShortUrl{}
	}

	if err == mongo.ErrNoDocuments {
		newShort := ShortUrl{
			ID:          str,
			Destination: url,
		}
		res, err := collection.InsertOne(ctx, newShort)
		if err != nil {
			log.Println("Insert failed:", err)
			return ShortUrl{}
		}
		log.Println("Short URL created with ID:", res.InsertedID)
		return newShort
	}

	log.Println("Short ID already exists, retrying not implemented here")
	return ShortUrl{}
}

func FindShortenedUrl(id string) *ShortUrl {
	shortUrls := GetShortenedUrls()

	for _, url := range shortUrls {
		if strings.Compare(url.ID, id) == 0 {
			return &url
		}
	}

	return nil
}
