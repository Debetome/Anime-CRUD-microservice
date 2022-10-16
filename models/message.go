package models

import (
	"io"
	"encoding/json"
)

type Message struct {
	Msg 	string 	`json: "msg,omitempty"`
	Status	string 	`json: "status,omitempty"`
}

func (p *Message) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Message) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}