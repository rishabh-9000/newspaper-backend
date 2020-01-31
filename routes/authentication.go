package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"newspaper-backend/config"
	"newspaper-backend/helper"
	"newspaper-backend/models"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Email : To get email where OTP will be sent
type Email struct {
	Email string `json:"email" bson:"email"`
}

// SendOTP : Sends OTP to User
func SendOTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var email Email
	var otpResponse models.OTP
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

	collecton := config.Client.Database(os.Getenv("db")).Collection("otp")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	otpResponse.Email = email.Email
	otpResponse.OTP = otp
	otpResponse.CreatedAt = time.Now()

	_, e = collecton.InsertOne(ctx, otpResponse)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")

	otpCollecton := config.Client.Database(os.Getenv("db")).Collection("otp")
	profileCollection := config.Client.Database(os.Getenv("db")).Collection("profile")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var otp models.OTP
	var profile models.Profile
	var profileExists models.Profile
	var finalResponse models.FinalResponse

	e := json.NewDecoder(r.Body).Decode(&otp)
	if e != nil {
		log.Println("Invalid Fields: ", e.Error())
		return
	}

	email := otp.Email
	otpCode := otp.OTP

	profile.Email = email

	_, e = otpCollecton.DeleteOne(ctx, bson.M{"email": email, "otp": otpCode})
	if e != nil {
		log.Println("Error in Removing: ", e.Error())
		return
	}

	_ = profileCollection.FindOne(ctx, bson.M{"email": email}).Decode(&profileExists)
	if profileExists.Email == "" {
		_, e = profileCollection.InsertOne(ctx, profile)
	}

	jwtToken, e := helper.EncodeJWT(email)
	if e != nil {
		log.Println("JWT Encoding Failed")
		return
	}

	finalResponse.Status = "success"
	finalResponse.Body = jwtToken

	json.NewEncoder(w).Encode(&finalResponse)
}

// GetUserEmail : Decodes JWT to Email
func GetUserEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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
