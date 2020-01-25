package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Profile : Model for user profile
type Profile struct {
	Name  string               `json:"name", bson:"name,omitempty"`
	Email string               `json:"email,omitempty" bson:"email,omitempty"`
	News  []primitive.ObjectID `json:"news" bson:"news,omitempty"`
}
