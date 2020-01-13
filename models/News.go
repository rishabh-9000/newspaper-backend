package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// News : News Model
type News struct {
	ID 		   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Host       string 			  `json:"host,omitempty" bson:"host,omitempty"`
	Category   string 			  `json:"category,omitempty" bson:"category,omitempty"`
	Headline   string 			  `json:"headline,omitempty" bson:"headline,omitempty"`
	Image 	   string			  `json:"image,omitempty" bson:"image,omitempty"`
	URL        string 			  `json:"url,omitempty" bson:"url,omitempty"`
	Date       string 			  `json:"date,omitempty" bson:"date,omitempty"`
	ClickCount int64              `json:"clickCount,omitempty" bson:"clickCount,omitempty"`
	Archived   bool               `json:"archived,omitempty" bson:"archived,omitempty"`
	CreatedAt  time.Time 		  `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
