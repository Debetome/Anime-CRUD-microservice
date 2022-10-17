package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/go-playground/validator"
	"encoding/json"
	"io"
)

type Anime struct {	
	ID 		  primitive.ObjectID  	`json:"id,omitempty"       bson:"_id,omitempty"`
	Emision   bool	  				`json:"emision"            bson:"emision"`
	Name 	  string  				`json:"name,omitempty"     bson:"name,omitempty"`
	Studio 	  string  				`json:"studio,omitempty"   bson:"studio,omitempty"`
	Watched   bool    				`json:"watched"   		   bson:"wacthed"`
	Seasons   int 	  				`json:"seasons,omitempty"  bson:"seasons,omitempty"`
}

type Animes []*Anime

func NewAnime(emision bool, name string, wacthed bool, seasons int) (a *Anime) {
	return &Anime{ Emision: emision, Name: name, Watched: wacthed, Seasons: seasons }
}

func (self *Anime) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(self)
}

func (self *Anime) Validate() error {
	validate := validator.New()
	return validate.Struct(self)
}