package models

type Personalidade struct {
	Nome    string `json:"nome"`
	Histori string `json:"historia"`
}

var Personalidades []Personalidade
