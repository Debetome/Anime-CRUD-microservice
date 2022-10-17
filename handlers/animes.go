package handlers

import (
	"github.com/gorilla/mux"
	"microservice/database"
	"microservice/models"
	"net/http"
	"log"
)

type Animes struct {
	database 	*database.Database
	logger		*log.Logger
}

func NewAnimes(database *database.Database, logger *log.Logger) *Animes {
	return &Animes{database, logger}
}

func (self *Animes) GetAnimes(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	animes, err := self.database.GetRecords()
	if err != nil {
		self.logger.Println(err)
		self.logger.Println("Unable to fetch anime records")
		http.Error(rw, "Unable to fetch anime records", http.StatusInternalServerError)
		return
	}

	err = animes.ToJSON(rw)
	if err != nil {
		self.logger.Println(err)
		self.logger.Println("Unable to parse response")
		http.Error(rw, "Bad request, unable to parse response", http.StatusBadRequest)
		return
	}

	self.logger.Println("Anime records fetched successfuly")
}

func (self *Animes) GetAnime(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	var anime *models.Anime
	anime, err := self.database.GetRecord(id)
	if err != nil {
		self.logger.Println(err)
		self.logger.Println("Record not found")
		http.Error(rw, "Record not found", http.StatusNotFound)
		return
	}

	err = anime.ToJSON(rw)
	if err != nil {
		self.logger.Println(err)
		self.logger.Println("Unable to parse response")
		http.Error(rw, "Bad request, unable to parse response", http.StatusBadRequest)
		return
	}

	self.logger.Printf("'%v' record fetched successfuly!", *anime)
}

func (self *Animes) PostAnime(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	newAnime := &models.Anime{}
	response := &models.Message{}
	msg := "Anime record added succesfuly!"

	err := newAnime.FromJSON(r.Body)
	if err != nil {
		self.logger.Println(r.Body)
		self.logger.Println(err)
		http.Error(rw, "Bad request, unable to parse body data", http.StatusBadRequest)
		return
	}

	err = self.database.AddRecord(newAnime)
	if err != nil {
		self.logger.Println("Unable to add record")
		http.Error(rw, "Bad request, unable to add record", http.StatusBadRequest)
		return
	}

	response.Msg = msg
	response.Status = "Success"

	err = response.ToJSON(rw)
	if err != nil {
		self.logger.Println(err)
		self.logger.Println("Unable to parse response")
		http.Error(rw, "Bad request, unable to parse response", http.StatusBadRequest)
		return
	}

	self.logger.Printf("'%v' record added successfuly!", *newAnime)
}

func (self *Animes) UpdateAnime(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	changedAnime := &models.Anime{}
	response := &models.Message{}
	msg := "Anime record updated successfuly!"

	err := changedAnime.FromJSON(r.Body)
	if err != nil {
		self.logger.Println(r.Body)
		self.logger.Println(err)
		http.Error(rw, "Bad request, unable to parse body data", http.StatusBadRequest)
		return
	}

	err = self.database.UpdateRecord(id, changedAnime)
	if err != nil {
		self.logger.Println(err)
		http.Error(rw, "Record not found", http.StatusNotFound)
		return
	}

	response.Msg = msg
	response.Status = "Success"

	err = response.ToJSON(rw)
	if err != nil {
		self.logger.Println(err)
		self.logger.Println("Unable to parse response")
		http.Error(rw, "Bad request, unable to parse response", http.StatusBadRequest)
		return
	}

	self.logger.Println("Record updated successfuly!")
}

func (self *Animes) DeleteAnime(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	response := &models.Message{}
	msg := "Anime record deleted successfuly!"

	err := self.database.DeleteRecord(id)
	if err != nil {
		self.logger.Println(err)
		self.logger.Println("Record does not exist!")
		http.Error(rw, "Record does not exist", http.StatusNotFound)
		return
	}

	response.Msg = msg
	response.Status = "Success"

	err = response.ToJSON(rw)
	if err != nil {
		self.logger.Println(err)
		self.logger.Println("Unable to parse response")
		http.Error(rw, "Bad request, unable to parse response", http.StatusBadRequest)
		return
	}

	self.logger.Println("Record deleted successfuly!")
}