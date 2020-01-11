package routes

import (
	"context"
	"net/http"
	"encoding/json"
	"newspaper-backend/config"
	"newspaper-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// AllNews : Returns all news
func AllNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	collection := config.Client.Database("newspaper").Collection("news")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var allNews []models.News
	var finalResponse models.FinalResponse

	cursor, e := collection.Find(ctx, bson.M{})
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + e.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var news models.News
		cursor.Decode(&news)
		allNews = append(allNews, news)
	}

	finalResponse.Status = "success"
	finalResponse.Body = allNews

	json.NewEncoder(w).Encode(finalResponse)
}

// BusinessNews : Returns all business news
func BusinessNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	collection := config.Client.Database("newspaper").Collection("news")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var allBusinessNews []models.News
	var finalResponse models.FinalResponse

	cursor, e := collection.Find(ctx, models.News{Category: "business"})
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + e.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var news models.News
		cursor.Decode(&news)
		allBusinessNews = append(allBusinessNews, news)
	}

	finalResponse.Status = "success"
	finalResponse.Body = allBusinessNews

	json.NewEncoder(w).Encode(finalResponse)
}

// SportsNews : Returns all sports news
func SportsNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	collection := config.Client.Database("newspaper").Collection("news")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var allSportsNews []models.News
	var finalResponse models.FinalResponse

	cursor, e := collection.Find(ctx, models.News{Category: "sports"})
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + e.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var news models.News
		cursor.Decode(&news)
		allSportsNews = append(allSportsNews, news)
	}
	
	finalResponse.Status = "success"
	finalResponse.Body = allSportsNews

	json.NewEncoder(w).Encode(finalResponse)
}

// EntertainmentNews : Returns all entertainment
func EntertainmentNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	collection := config.Client.Database("newspaper").Collection("news")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var allEntertainmentNews []models.News
	var finalResponse models.FinalResponse

	cursor, e := collection.Find(ctx, models.News{Category: "entertainment"})
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + e.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var news models.News
		cursor.Decode(&news)
		allEntertainmentNews = append(allEntertainmentNews, news)
	}

	finalResponse.Status = "success"
	finalResponse.Body = allEntertainmentNews

	json.NewEncoder(w).Encode(finalResponse)
}
