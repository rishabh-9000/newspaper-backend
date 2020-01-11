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

	json.NewEncoder(w).Encode(allNews)
}
