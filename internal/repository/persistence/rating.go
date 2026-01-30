package persistence

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rating struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     string             `bson:"user_id"`
	ProductID  string             `bson:"product_id"`
	Type       string             `bson:"type"`
	Value      int                `bson:"value"`
	InsertedAt time.Time          `bson:"inserted_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}
