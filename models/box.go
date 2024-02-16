package models

type Content struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type Box struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Content []Content `json:"content"`
}
