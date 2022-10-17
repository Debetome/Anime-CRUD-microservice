package models

import (
	"io"
	"encoding/json"
)

type Record struct {
	Msg		string	`json:"msg,omitempty"`
	Status	string 	`json:"status,omitempty"`
	Anime	*Anime	`json:"anime"`
}

func (p *Record) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}