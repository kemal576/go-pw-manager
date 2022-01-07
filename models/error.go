package models

type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
