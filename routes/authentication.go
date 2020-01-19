package routes

import (
	"log"
	"time"
	"context"
	"net/http"
	"encoding/json"
	"newspaper-backend/config"
	"newspaper-backend/helper"
	"newspaper-backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

type Email struct {
	Email string `json:"email" bson:"email"`
}

type OTP struct {
	Email 	  string 	`json:"email" bson:"email"`
	OTP   	  string 	`json:"otp" bson:"otp"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

// SendOTP : Sends OTP to User
func SendOTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	collecton := config.Client.Database("newspaper").Collection("otp")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var email Email
	var otpResponse OTP
	var finalResponse models.FinalResponse

	e := json.NewDecoder(r.Body).Decode(&email)
	if e != nil {
		log.Println("Requires Email: ", e.Error())
		return
	}

	otp, e := helper.GenerateOTP(6)
	if e != nil {
		log.Println("OTP generation failed: ", e.Error())
		return
	}

	emailSubject := "One Time Password for Your Newspaper"
	emailBody := "You OTP is: " + otp

	result, e := helper.SendEmail(email.Email, emailSubject, emailBody)
	if e != nil {
		log.Println("Failed to send email: ", e.Error())
		return
	}

	otpResponse.Email = email.Email
	otpResponse.OTP = otp
	otpResponse.CreatedAt = time.Now()

	_, e = collecton.InsertOne(ctx, otpResponse,)
	if e != nil {
		log.Println("Failed to enter in DB: ", e.Error())
		return
	}

	finalResponse.Status = "success"
	finalResponse.Body = result

	json.NewEncoder(w).Encode(finalResponse)
}

// Authenticate : Takes Email and OTP and authenticate
func Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	collecton := config.Client.Database("newspaper").Collection("otp")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var otp OTP
	var finalResponse models.FinalResponse

	e := json.NewDecoder(r.Body).Decode(&otp)
	if e != nil {
		log.Println("Invalid Fields: ", e.Error())
		return
	}

	email := otp.Email
	otpCode := otp.OTP

	_, e = collecton.DeleteOne(ctx, bson.M{"email": email, "otp": otpCode})
	if e != nil {
		log.Println("Error in Removing: ", e.Error())
		return
	}

	jwtToken, e := helper.EncodeJWT(email)

	finalResponse.Status = "success"
	finalResponse.Body = jwtToken

	json.NewEncoder(w).Encode(&finalResponse)
}

// GetUserEmail : Decodes JWT to Email
func GetUserEmail(w http.ResponseWriter, r *http.Request) {
	var finalResponse models.FinalResponse	

	tokenString := r.Header["X-Auth-Token"][0]

	email, e := helper.DecodeJWT(tokenString)
	if e != nil {
		log.Println("Failed in Decoding")
		return
	}

	finalResponse.Status = "success"
	finalResponse.Body = email

	json.NewEncoder(w).Encode(&finalResponse)
}