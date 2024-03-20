package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"foodname" validate:"required,min=2max=100"`
	Price      *float64           `json:"foodprice" validate:"required"`
	Food_Image  *string            `json:foodimage validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`

	Food_id string  `json:"food_id"`
	Menu_id *string `json:"menu_id"`
}
