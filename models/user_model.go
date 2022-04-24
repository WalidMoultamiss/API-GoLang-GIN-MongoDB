package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Name    string             `json:"name,omitempty" validate:"required"`
	Code    string             `json:"code,omitempty" validate:"required"`
	Balance string             `json:"balance,omitempty" validate:"required"`
}
