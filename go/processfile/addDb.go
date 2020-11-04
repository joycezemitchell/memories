package gd

import (
	"fmt"

	"github.com/joho/godotenv"
		
	config "allyapps.com/memories/config"
	models "allyapps.com/memories/models"

	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddToDb file
func AddToDb(data map[string]string) {

	godotenv.Load("/var/www/memories.allyapps.com.private/memories.env")
	// https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go
	// https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
	// call api to add to database

	fmt.Println("Add to db")

	// fmt.Println(data["thumbnail"])

	/*requestBody, _ := json.Marshal(map[string]string{
		"title":       "Australia Day 1 trip ddddd",
		"description": "Our first trip in Sydney, Australia",
		"video":       "index.hml",
		"thumbnail":   "https://material-ui.com/static/images/grid-list/vegetables.jpg",
		"created_at":  "324234324",
		"updated_at":  "324234",
	})*/


	// I think I will need to do this directly rather than using that stupid grpc thing
	
	// Disabled http post  add DB because there is now an additional security middleware and I am too tired to work on it. 
	
	/*requestBody, _ := json.Marshal(map[string]string{
		"title":       data["title"],
		"description": data["description"],
		"video":       data["video"],
		"thumbnail":   data["thumbnail"],
		"created_at":  data["created_at"],
		"updated_at":  data["updated_at"],
	})

	resp, _ := http.Post("http://localhost:13000/api/v1/videos", "application/json", bytes.NewBuffer(requestBody))
	resp, _ := http.Post("http://"+os.Getenv("APIURL")+":"+os.Getenv("APIPORT")+"/api/v1/videos", "application/json", bytes.NewBuffer(requestBody))
    defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(string(body)) */

	// Add to DB the old skul way
	
	dataSource := models.Videos{
		Title:       data["title"],
		Description: data["description"],
		Video:       data["video"],
		Thumbnail:   data["thumbnail"],
		CreatedAt:   data["created_at"],
		UpdatedAt:   data["updated_at"],
	}

	res, _ := config.Database.Collection("videos").InsertOne(context.Background(), dataSource)
	oid, _ := res.InsertedID.(primitive.ObjectID)

	fmt.Println(oid)

}
