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
		http.Error(rw, "Unable to fetch anime records", http.StatusInternalServerError)
		return
	}

	err = animes.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (self *Animes) GetAnime(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
}

func (self *Animes) PostAnime(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	newAnime := &models.Anime{}

	err := newAnime.FromJSON(r.Body)
	if err != nil {
		self.logger.Println(r.Body)
		self.logger.Println(err)
		http.Error(rw, "Unable to parse data arguments", http.StatusBadRequest)
	}

	self.database.AddRecord(newAnime)
	self.logger.Println("Anime record created succesfuly!")
}

func (self *Animes) UpdateAnime(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	var _ models.Anime
	_ = mux.Vars(r)

}

func (self *Animes) DeleteAnime(rw http.ResponseWriter, r *http.Request) {

}