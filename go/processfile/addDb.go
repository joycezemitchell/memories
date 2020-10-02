package gd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// AddToDb file
func AddToDb(data map[string]string) {

	godotenv.Load("/var/www/memories.allyapps.com/memories.env")
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

	requestBody, _ := json.Marshal(map[string]string{
		"title":       data["title"],
		"description": data["description"],
		"video":       data["video"],
		"thumbnail":   data["thumbnail"],
		"created_at":  data["created_at"],
		"updated_at":  data["updated_at"],
	})

	//resp, _ := http.Post("http://localhost:13000/api/v1/videos", "application/json", bytes.NewBuffer(requestBody))
	resp, _ := http.Post("http://"+os.Getenv("APIURL")+":"+os.Getenv("APIPORT")+"/api/v1/videos", "application/json", bytes.NewBuffer(requestBody))

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(string(body))
}
