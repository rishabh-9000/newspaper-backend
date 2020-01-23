package models

// FinalResponse : Model for final response
type FinalResponse struct {
	Status string      `json:"status,omitempty" bson:"status,omitempty"`
	Body   interface{} `json:"body,omitempty" bson:"body,omitempty"`
}
