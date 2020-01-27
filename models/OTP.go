package models

import "time"

// OTP : Model for OTP
type OTP struct {
	Email     string    `json:"email" bson:"email"`
	OTP       string    `json:"otp" bson:"otp"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
