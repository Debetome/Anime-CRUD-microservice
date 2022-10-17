package models

import (
	"io"
	"encoding/json"
)

type Records struct {
	Msg     string  `json:"msg,omitempty"`
	Status	string 	`json:"status,omitempty"`
	Animes	Animes	`json:"animes"`
}

func (p *Records) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}