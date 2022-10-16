package tests

import (
	"microservice/database"
	"microservice/handlers"
	"testing"
	"log"
)

func TestFetchRecords(t *testing.T) {
	logger := log.New(os.Stdout, "anime-api ", log.LstdFlags)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTION_STRING))
	if err != nil {
		logger.Fatal(err)
	}

	defer client.Disconnect(ctx)

	database := database.NewDatabase(&client, &logger)
	records, err := database.GetRecords()
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log("Database instance created successfuly")
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