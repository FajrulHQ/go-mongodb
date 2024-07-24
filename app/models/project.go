package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Projects struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Link        string             `json:"link"`
	Start       *time.Time         `json:"start"`
	End         *time.Time         `json:"end"`
	Created     time.Time          `json:"created"`
	Updated     time.Time          `json:"updated"`
}

type ProjectRequest struct {
	Name        string
	Description string
	Link        string
	Start       *time.Time
	End         *time.Time
}
