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

	log.Fatal(http.ListenAndServe(":5000", router))
}

// Test : Test Route
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Running")
}
