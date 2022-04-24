package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Facture struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Serial  string             `json:"serial,omitempty"`
	Price   int                `json:"price,omitempty"`
	Company string             `json:"company,omitempty"`
	Status  string             `json:"status,omitempty"`
}
