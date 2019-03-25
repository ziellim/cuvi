package model

import (
	"encoding/json"
)

type CV struct {
	FirstName string  `json:"firstname"`
	LastName  string  `json:"lastname"`
	Title     string  `json:"title"`
	Contact   Contact `json:"contact"`
}

func LoadCV() *CV {
	return &CV{
		LastName:  "Larson",
		FirstName: "Nicky",
		Title:     "City Hunter",
		Contact: Contact{
			Email:     "nickyson@gmail.com",
			Telephone: " +81 3 0029 7548",
			Location:  "107-8503, Japon",
		},
	}
}

func (cv *CV) Marshal() ([]byte, error) {
	return json.Marshal(cv)
}
