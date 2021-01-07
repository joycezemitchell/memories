package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	gd "allyapps.com/memories/processfile"
)

func main() {

	sourceDirectory := "/mnt/e/Project/memories/go/fileuploader/temp_mp4"
	targetDirecotry := "/mnt/e/Project/memories/go/fileuploader/temp_hls"

	sd := flag.String("source-dir", sourceDirectory, "Source Directory")
	td := flag.String("target-dir", targetDirecotry, "Target Directory")

	flag.Parse()

	fmt.Println("Grabing all files from:", *sd, td)

	//Reference: https://flaviocopes.com/go-list-files/

	//	- iterate to all mp4 files in temp_mp4 directory
	filepath.Walk(*sd, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		// grab the date
		date := gd.GetDate(info)

		// grab the directory name
		directoryName := gd.GenerateDirectoryNameFromFilename(info.Name())

		// Process the file (generate videos, grab thumbnail, upload to digital space )
		processMp4(info.Name(), *sd, *td, directoryName)

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
		// https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go
		// call api to add to database

		gd.AddToDb(data)

		// Delete file after uploading
		os.Remove(sourceDirectory + "/" + info.Name())
		os.Remove(targetDirecotry + "/" + info.Name())

		return nil
	})

}

func processMp4(file string, sourceDirectory string, targetDirectory string, directoryName string) {
	fmt.Println("Processing", file)

	cmd := exec.Command("cp", sourceDirectory+"/"+file, targetDirectory+"/"+file)
	cmd.Stdout = os.Stdout
	cmd.Run()

	// mv mp4 file to temp_hls directory
	fmt.Println("Moving " + file + " to " + targetDirectory + "...")

	// Check video rotation. Check if portrait

	// convert mp4 to hls
	gd.ConvertToHls(file, targetDirectory)

	// - grab thumbnail
	gd.GenerateThumbnails(file, targetDirectory)

	// iterate all files inside temp_hls directory and upload it to DO
	filepath.Walk(targetDirectory, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".mp4" || filepath.Ext(path) == ".MOV" {
			return nil
		}

		// Upload to digital space
		gd.Upload(info.Name(), targetDirectory, directoryName)

		// Delete file after uploading
		os.Remove(targetDirectory + "/" + info.Name())

		return nil
	})

}
