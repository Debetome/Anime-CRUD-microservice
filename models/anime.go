package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"encoding/json"
	"io"
)

type Anime struct {	
	ID 		  primitive.ObjectID  	`json: "id, omitempty"       bson: "_id, omitempty"`
	Emision   bool	  				`json: "emision, omitempty"  bson: "emision, omitempty"`
	Name 	  string  				`json: "name, omitempty"     bson: "name, omitempty"`		
	Watched   bool    				`json: "watched, omitempty"  bson: "wacthed, omitempty"`
	Seasons   int 	  				`json: "seasons, omitempty"  bson: "seasons, omitempty"`
}

type Animes []*Anime

func NewAnime(emission bool, name string, wacthed bool, seasons int) (a *Anime) {
	return &Anime{ Emision: emision, Name: name, Watched: wacthed, Seasons: seasons }
}

func (p *Anime) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Animes) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}