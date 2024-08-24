package models

import (
	"time"
)
type Recipe struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Ingredients     []string  `json:"ingredients"`
	Instructions    string    `json:"instructions"`
	PreparationTime int       `json:"preparationTime"`
	CreationDate    time.Time `json:"creationDate"`
}
