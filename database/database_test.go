package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"microservice/models"
	"context"
	"testing"
	"reflect"
	"log"
	"os"
	"time"
)

const CONNECTION_STRING = "mongodb://pepeluis:12345678@localhost:27017/?authSource=People"

func TestFetchRecords(t *testing.T) {
	logger := log.New(os.Stdout, "test-anime-api ", log.LstdFlags)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTION_STRING))
	if err != nil {
		logger.Fatal(err)
	}

	defer client.Disconnect(ctx)

	database := NewDatabase(client, logger)
	records, err := database.GetRecords()

	if err != nil {
		t.Error(err)
		t.Fail()
	}
	
	if len(records) == 0 {
		t.Error("Either database struct was not able to fetch any records OR none have been saved yet!")
		t.Fail()
	}

	for _, record := range records {
		if reflect.TypeOf(record) != reflect.TypeOf(&models.Anime{}) {
			t.Error("Record type do not match!")
			t.Fail()
		}
	}
}

func TestFetchRecord(t *testing.T) {
	logger := log.New(os.Stdout, "anime-api ", log.LstdFlags)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTION_STRING))
	if err != nil {
		logger.Fatal(err)
	}

	defer client.Disconnect(ctx)

	// Change this according to database or collection you test this with
	const TEST_ID = "634cc5de7962de5d3e206f86"

	database := NewDatabase(client, logger)
	record, err := database.GetRecord(TEST_ID)

	if err != nil {
		t.Error(err)
		t.Error("Unable to fetch record or record does not exists ...")
		t.Fail()
	}

	if reflect.TypeOf(record) != reflect.TypeOf(&models.Anime{}) {
		t.Error("Record type does not match! ...")
		t.Fail()
	}
}

func TestUpdateRecord(t *testing.T) {
	logger := log.New(os.Stdout, "anime-api ", log.LstdFlags)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTION_STRING))
	if err != nil {
		logger.Fatal(err)
	}

	defer client.Disconnect(ctx)
}