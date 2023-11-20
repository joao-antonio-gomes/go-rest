package models

type Persona struct {
	Name      string `json:"name"`
	Biography string `json:"biography"`
}

var Personas []Persona
