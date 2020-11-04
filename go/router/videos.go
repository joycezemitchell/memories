package router

import (
	"context"
	"encoding/json"
	"net/http"

	config "allyapps.com/memories/config"
	models "allyapps.com/memories/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Videos
func Videos() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		videos := grabAll()
		json.NewEncoder(w).Encode(videos)

	})
}

func grabAll() []models.Videos {
	pipeline := []bson.D{bson.D{{"$sample", bson.D{{"size", 500}}}}}
	cur, _ := config.Database.Collection("videos").Aggregate(context.Background(), pipeline)

	var videos []models.Videos

	for cur.Next(context.TODO()) {
		data := models.Videos{}
		cur.Decode(&data)
		videos = append(videos, data)
	}

	return videos
}
