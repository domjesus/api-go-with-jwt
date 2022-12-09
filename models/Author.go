package models

type Author struct {
	ID     int
	Name   string `json:"name"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
}

var Authors []Author
