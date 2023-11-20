package models

type Persona struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}

var Personas []Persona
