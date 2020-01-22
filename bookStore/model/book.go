package model

type Book struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int64  `json:"quantity"`
}
