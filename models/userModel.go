package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	First_name string             `json:"first_name" validate:"required,min=2,max=100"`
	Last_name  string             `json:"last_name" validate:"required,min=2,max=100"`
	Email      string             `json:"email" validate:"email,required"`
	Phone      string             `json:"phone" validate:"required"`
	// omitempty == if no data 
	Created_at time.Time          `json:"created_at,omitempty"`
	Updated_at time.Time          `json:"updated_at,omitempty"`
}
