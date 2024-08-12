package models

type Message struct {
	Text  string `json:"text"`
	Class string `json:"class"`
	Alt   string `json:"alt"`
}
