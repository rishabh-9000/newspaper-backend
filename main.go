package main

import (
	"fmt"
	"log"
	"net/http"
	"newspaper-backend/config"
	"newspaper-backend/routes"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Application Started")

	config.DbConnect()

	router := mux.NewRouter()

	// Test Route
	router.HandleFunc(
		"/",
		Test).Methods("GET")

	router.HandleFunc(
		"/api/all-news",
		routes.AllNews).Methods("GET")

	router.HandleFunc(
		"/api/business-news",
		routes.BusinessNews).Methods("GET")

	router.HandleFunc(
		"/api/sports-news",
		routes.SportsNews).Methods("GET")

	router.HandleFunc(
		"/api/entertainment-news",
		routes.EntertainmentNews).Methods("GET")

	router.HandleFunc(
		"/api/click-count/{news_id}",
		routes.ClickCount).Methods("GET")

	router.HandleFunc(
		"/api/most-viewed-news",
		routes.MostViewedNews).Methods("GET")

	router.HandleFunc(
		"/api/host-list",
		routes.GetHosts).Methods("GET")

	router.HandleFunc(
		"/api/send-otp",
		routes.SendOTP).Methods("POST")

	router.HandleFunc(
		"/api/authenticate",
		routes.Authenticate).Methods("POST")

	router.HandleFunc(
		"/api/get-user",
		routes.GetUserEmail).Methods("GET")

	router.HandleFunc(
		"/api/save-news/{id}",
		routes.SaveNews).Methods("GET")

	router.HandleFunc(
		"/api/saved-news",
		routes.GetSavedNews).Methods("GET")

	router.HandleFunc(
		"/api/remove-news/{id}",
		routes.RemoveSavedNews).Methods("GET")

	router.HandleFunc(
		"/api/user",
		routes.GetUser).Methods("GET")

	router.HandleFunc(
		"/api/update-user",
		routes.UpdateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", router))
}

// Test : Test Route
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Running")
}
