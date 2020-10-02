package main

import (
	gd "allyapps.com/memories/processfile"
)

func main() {

	data := map[string]string{
		"title":       "Australia Day 2 trip",
		"description": "Our first trip in Sydney, Australia",
		"video":       "index.hml",
		"thumbnail":   "https://material-ui.com/static/images/grid-list/vegetables.jpg",
		"created_at":  "324234324",
		"updated_at":  "324234",
	}

	gd.AddToDb(data)

	fmt.Println("test")


}
