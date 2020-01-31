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

type userDetail struct {
	Name  string
	Email string
}

// GetUser : Gets user data
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	collection := config.Client.Database(os.Getenv("db")).Collection("profile")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var profile models.Profile
	var finalResponse models.FinalResponse

	JWT := r.Header["X-Auth-Token"][0]
	email, e := helper.DecodeJWT(JWT)
	if e != nil {
		log.Println("Unauthorized")

		finalResponse.Status = "failed"
		finalResponse.Body = "unauthorized"

		json.NewEncoder(w).Encode(finalResponse)
		return
	}

	e = collection.FindOne(ctx, bson.M{"email": email}).Decode(&profile)
	if e != nil {
		log.Println("Didn't find record in DB")
		return
	}

	finalResponse.Status = "success"
	finalResponse.Body = profile

	json.NewEncoder(w).Encode(finalResponse)
}

// UpdateUser : Update user detail
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	collection := config.Client.Database(os.Getenv("db")).Collection("profile")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user userDetail
	var finalResponse models.FinalResponse

	JWT := r.Header["X-Auth-Token"][0]
	userEmail, e := helper.DecodeJWT(JWT)
	if e != nil {
		log.Println("Unauthorized")

		finalResponse.Status = "status"
		finalResponse.Body = "unauthorized"

		json.NewEncoder(w).Encode(finalResponse)
		return
	}
	e = json.NewDecoder(r.Body).Decode(&user)
	if e != nil {
		log.Println("Name is required")
		return
	}

	filter := bson.M{"email": userEmail}
	update := bson.M{"$set": bson.M{"name": user.Name, "email": user.Email}}
	result, e := collection.UpdateOne(ctx, filter, update)
	if e != nil {
		log.Println("Failed to update in Mongo")
		return
	}

	finalResponse.Status = "success"
	finalResponse.Body = result

	json.NewEncoder(w).Encode(finalResponse)
}
