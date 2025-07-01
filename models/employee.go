package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FullName  string             `bson:"full_name" json:"full_name"`
	Email     string             `bson:"email" json:"email"`
	Position  string             `bson:"position" json:"position"`
	Salary    float64            `bson:"salary" json:"salary"`
	IsActive  bool               `bson:"is_active" json:"is_active"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
