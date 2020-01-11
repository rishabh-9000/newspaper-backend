package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// News : News Model
type News struct {
	ID 		 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Host     string 			`json:"host,omitempty" bson:"host,omitempty"`
	Category string 			`json:"category,omitempty" bson:"category,omitempty"`
	Headline string 			`json:"headline,omitempty" bson:"headline,omitempty"`
	URL      string 			`json:"url,omitempty" bson:"url,omitempty"`
	Date     string 			`json:"date,omitempty" bson:"date,omitempty"`
}
