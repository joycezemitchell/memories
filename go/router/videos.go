package router

import (
	"fmt"
	"strconv"
	"encoding/json"
	"net/http"
	. "github.com/gobeam/mongo-go-pagination"
	config "allyapps.com/memories/config"
	models "allyapps.com/memories/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Videos
func Videos() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		pages, _ := r.URL.Query()["p"]
		p := pages[0]
		fmt.Println("Current Page:" + string(p))
		px,_ := strconv.ParseInt(p, 10, 64)
		videos := grabAll(px)
		json.NewEncoder(w).Encode(videos)
	})
}

func grabAll(p int64) []models.Videos {
	// pipeline := []bson.D{bson.D{{"$sample", bson.D{{"size", 30}}}}}
	// cur, _ := config.Database.Collection("videos").Aggregate(context.Background(), pipeline)
	// cur, _ := config.Database.Collection("videos").Find(context.TODO(), bson.M{})

	/*for cur.Next(context.TODO()) {
		data := models.Videos{}
		cur.Decode(&data)
		videos = append(videos, data)
	}*/


	// Using pagination - https://github.com/gobeam/mongo-go-pagination

	var videos []models.Videos

	collection := config.Database.Collection("videos")
	filter := bson.M{}

	paginatedData, err := New(collection).Limit(30).Page(p).Filter(filter).Find()
	if err != nil {
		panic(err)
	}
	 
	for _, raw := range paginatedData.Data {
		data := models.Videos{}
		if marshallErr := bson.Unmarshal(raw, &data); marshallErr == nil {
			videos = append(videos, data)
		}
	}
		
	return videos
}
