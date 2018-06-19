package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"time"
	"os"
	"github.com/rhymond/go-money"
	"errors"
	"dexmohq.com/dexpenses-classify/internal/pkg/models"
)

var config *mongo.Collection
var receipts *mongo.Collection

func init() {
	mongoUri := os.Getenv("MONGO_URI")
	if len(mongoUri) == 0 {
		log.Fatal("MONGO_URI environment variable is missing")
		return
	}
	client, err := mongo.NewClient(mongoUri)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return
	}
	db := client.Database("dexmohq")
	config = db.Collection("config")
	receipts = db.Collection("receipts")
	//cursor, err := config.Find(context.Background(), nil)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//defer cursor.Close(context.Background())
	//for cursor.Next(context.Background()) {
	//	document := bson.NewDocument()
	//	err := cursor.Decode(document)
	//	if err != nil {
	//		log.Fatal(err)
	//		continue
	//	}
	//
	//}
}

func HandleRequest(ctx context.Context, e events.DynamoDBEvent) (string, error) {
	if config == nil || receipts == nil {
		message := "could not grab collection"
		log.Fatal(message)
		return "Error", errors.New(message)
	}
	//log.Print(e.Records)
	//for _, record := range e.Records {
	//	fmt.Printf("Processing data for event id %s, type %s.\n", record.EventID, record.EventName)
	//	println(record.Change.NewImage)
	//}
	result, err := receipts.InsertOne(context.Background(),
		dexpenses.Receipt{ID: objectid.New(), Date: time.Now(), Time: time.Now(),
			Total: dexpenses.AsPersistentMoney(money.New(2999, "EUR")),
			PaymentMethod: dexpenses.Cash, Category: "food"})
	if err != nil {
		log.Fatal(err)
		return "Error", err
	}
	log.Printf("Inserted receipt %s", result.InsertedID)
	return "Success", nil
}

func main() {
	lambda.Start(HandleRequest)
}
