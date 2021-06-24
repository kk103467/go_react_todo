package model

type Todo struct {
	Id          int    `json:"id"`
	Text        string `json:"text"`
	IsCompleted bool   `json:"isCompleted"`
}
