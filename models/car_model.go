package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Car struct {
	Id            primitive.ObjectID `json:"id,omitempty"`
	Matriculation string             `json:"matriculation,omitempty"`
	Year          int                `json:"year,omitempty"`
	HorsePower    int                `json:"horsePower,omitempty"`
	Gas           string             `json:"gas,omitempty"`
	Status        string             `json:"status,omitempty"`
}
