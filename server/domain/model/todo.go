package model

type Todo struct {
	Id        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}
