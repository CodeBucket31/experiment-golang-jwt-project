package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	First_name *string            `json:"first_name" validate:"required,min=2,max=100" `
	Last_name  *string            `json:"last_name" validate:"required,min=2,max=100"`
	Passwoed   *string            `json:"Password" validate:"required,min=6"`
	Email      *string            `json:"email" validate:"required,email"`
	Phone      string             `json:"phone" validate:"required"` // or `binding:"required"`
	// Email      *string            `json:"email" validate:"email,required"`
	// Phone      string             `json:"phone" binding:"required"`
	// Phone         *string            `json:"phone" validate:"requird"`
	Token         *string   `json:"token"`
	User_Type     *string   `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token *string   `json:"refresh_token"`
	Created_at    time.Time `josn:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	User_id       string    `json:"user_id"`
}
