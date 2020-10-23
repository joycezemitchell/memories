package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	gd "allyapps.com/memories/processfile"
)

func main() {
	http.HandleFunc("/", postCreate)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	http.ListenAndServe(":2030", nil)
}

func postCreate(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var header = r.Header.Get("Hook-Name")
	header = strings.TrimSpace(header)

	var d map[string]map[string]map[string]interface{}

	json.Unmarshal([]byte(body), &d)
	// ioutil.WriteFile("test101.txt", body, 0644)

	// - parse json - parse filename
	if header == "post-finish" {

		st := fmt.Sprintf("%v", d["Upload"]["Storage"]["Path"])
		s := strings.Split(st, "/")
		s1 := s[:len(s)-1]
		s3 := s[:len(s)-2]

		originalFilename := fmt.Sprintf("%v", d["Upload"]["MetaData"]["filename"])
		tusFilename := s[len(s)-1]                // storage filename
		tusStorage := strings.Join(s1, "/")       // storage directory
		tusTempDirectory := strings.Join(s3, "/") // temp hls directory

		// - rename file

		fmt.Println(r.Header)
		fmt.Println("---------------------------")
		fmt.Println(header)
		fmt.Println(originalFilename)
		fmt.Println(tusFilename)
		fmt.Println(tusStorage)
		fmt.Println(tusTempDirectory)

		os.Chdir(tusStorage)

		// Copy file to hls and rename it to original filename
		cmd := exec.Command("mv", tusFilename, "hls/"+originalFilename)
		cmd.Stdout = os.Stdout
		cmd.Run()

		// Delete .info file
		os.Remove(tusStorage + "/" + tusFilename + ".info")

		// grab the directory name
		directoryName := gd.GenerateDirectoryNameFromFilename(originalFilename)

		// convert mp4 to hls
		gd.ConvertToHls(originalFilename, tusStorage+"/hls/")

		// - grab thumbnail
		gd.GenerateThumbnails(originalFilename, tusStorage+"/hls/")

		var date string

		// iterate all files inside temp_hls directory and upload it to DO
		filepath.Walk(tusStorage+"/hls/", func(path string, info os.FileInfo, err error) error {

			if info.IsDir() {
				return nil
			}

			if filepath.Ext(path) == ".mp4" || filepath.Ext(path) == ".MOV" {
				// grab the date
				date = gd.GetDate(info)
				return nil
			}

			// Upload to digital space
			gd.Upload(info.Name(), tusStorage+"/hls/", directoryName)

			// Delete file after uploading
			os.Remove(tusStorage + "/hls/" + info.Name())

			return nil
		})

		// Add to database
		data := map[string]string{
			"title":       "",
			"description": "",
			"video":       directoryName,
			"thumbnail":   "thumb.jpg",
			"created_at":  date,
			"updated_at":  "",
		}

		fmt.Println(data)

		// Add to DB
		gd.AddToDb(data)
		os.Remove(tusStorage + "/hls/" + originalFilename)
	}

	// s3cmd ls s3://allyapps
}
