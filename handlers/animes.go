package handlers

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"microservice/database"
	"microservice/models"
	"net/http"
	"log"
)

type Animes struct {
	database 	*database.Database
	logger		*log.Logger
}

func NewAnimes(database *database.Database, logger *log.Logger) *Person {
	return &Animes{database, logger}
}

func (self *Animes) GetAnimes(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	animes := self.database.GetRecords(rw)
	err := animes.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
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

	self.database.AddRecord(&newAnime)
	self.logger.Println("Anime record created succesfuly!")
}

func (self *Animes) UpdateAnime(rw http.ResponseWriter, r *http.Request) {+
	rw.Header().Add("Content-Type", "application/json")

	var anime Anime
	params = mux.Vars(r)
	id, err = primitive.ObjectIDFromHex(params["id"])


}

func (self *Animes) DeleteAnime(rw http.ResponseWriter, r *http.Request) {

}