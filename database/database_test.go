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

// Change this according to your case
const TEST_ID = "634cda607229e5428cafec5f"
const TEST_REMOVE_ID = "634cdaeed02bf52ced287121"


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

	database := NewDatabase(client, logger)
	record, err := database.GetRecord(TEST_ID)

	if err != nil {
		t.Error(err)
		t.Error("Unable to fetch record or record does not exists ...")
		t.Fail()
	}

	if reflect.TypeOf(record) != reflect.TypeOf(&models.Anime{}) {
		t.Error(err)
		t.Error("Record type does not match! ...")
		t.Fail()
	}
}

func TestAddRecord(t *testing.T) {
	logger := log.New(os.Stdout, "anime-api ", log.LstdFlags)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTION_STRING))
	if err != nil {
		logger.Fatal(err)
	}

	defer client.Disconnect(ctx)

	database := NewDatabase(client, logger)
	newAnime := models.NewAnime(false, "Bleach", false, 2)
	err = database.AddRecord(newAnime)

	if err != nil {
		t.Error(err)
		t.Error("Unable to save anime record! ...")
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

	database := NewDatabase(client, logger)
	err = database.UpdateRecord(TEST_ID, &models.Anime{Name: "Shingeki no Kyojin"})
	
	if err != nil {
		t.Error(err)
		t.Error("Unable to update record! ...")
		t.Fail()
	}
}

func TestDeleteRecord(t *testing.T) {
	logger := log.New(os.Stdout, "anime-api ", log.LstdFlags)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTION_STRING))
	if err != nil {
		logger.Fatal(err)
	}

	defer client.Disconnect(ctx)

	database := NewDatabase(client, logger)
	err = database.DeleteRecord(TEST_REMOVE_ID)
	if err != nil {
		t.Error(err)
		t.Error("Unable to delete record! ...")
		t.Fail()
	}
}