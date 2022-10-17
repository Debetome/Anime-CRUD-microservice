package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"microservice/models"
	"context"
	"log"
	"time"
)

type Database struct {
	client 	  *mongo.Client
	logger 	  *log.Logger
}

func NewDatabase(client *mongo.Client, logger *log.Logger) *Database {
	return &Database{client, logger}
}

func (self *Database) GetRecords() (animeRecords models.Animes, err error){	
	collection := self.client.Database("AnimesDB").Collection("animeList")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	
	if err != nil {
		self.logger.Println(err)
		self.logger.Println("Unable to get records!")
		return
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var anime models.Anime
		cursor.Decode(&anime)
		animeRecords = append(animeRecords, &anime)
	}

	if err = cursor.Err(); err != nil {
		self.logger.Println(err)
		self.logger.Println("Unable to get records!")
		return
	}

	return
}

func (self *Database) GetRecord(id string) (record *models.Anime, err error) {
	collection := self.client.Database("AnimesDB").Collection("animeList")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var hexID primitive.ObjectID
	hexID, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		self.logger.Println(err)
		self.logger.Println("Not valid Hex ID for getting record!")
		return
	}

	err = collection.FindOne(ctx, bson.M{"_id": hexID}).Decode(&record)
	if err != nil {
		self.logger.Println(err)
		self.logger.Printf("Record does not exists in this collection!")
		return
	}

	return
}

func (self *Database) AddRecord(anime *models.Anime) (err error) {
	collection := self.client.Database("AnimesDB").Collection("animeList")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = collection.InsertOne(
		ctx, 
		bson.M{"emision": anime.Emision, "name": anime.Name, "watched": anime.Watched, "seasons": anime.Seasons},
	)

	if err != nil {
		self.logger.Println(err)
		self.logger.Println("Unable to insert record")
		return
	}

	return
}

func (self *Database) UpdateRecord(id string, changed *models.Anime) (err error) {
	_, err = self.GetRecord(id)
	if err != nil {
		self.logger.Println(err)
		self.logger.Println("Unable to get record for updating")
		return
	}
	
	collection := self.client.Database("AnimesDB").Collection("animeList")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	var hexID primitive.ObjectID
	hexID, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		self.logger.Println(err)
		self.logger.Println("Not valid Hex ID for updating record!")
		return
	}

	_, err = collection.UpdateOne(
		ctx, 
		bson.M{"_id": hexID}, 
		bson.M{"$set": changed},
	)

	if err != nil {
		self.logger.Println(err)
		self.logger.Println("Unable to update record!")
		return
	}

	return
}

func (self *Database) DeleteRecord() {
	self.logger.Println(self.client)
}