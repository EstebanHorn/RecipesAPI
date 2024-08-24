package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name            string             `json:"name"`
	Ingredients     []string           `json:"ingredients"`
	Instructions    string             `json:"instructions"`
	PreparationTime int                `json:"preparationTime"`
	CreationDate    time.Time          `json:"creationDate"`
}
