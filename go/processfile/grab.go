package gd

import (
	"fmt"
	"os"
	"os/exec"
)

// GenerateThumbnails file
func GenerateThumbnails(file string, targetDirecotry string) {
	fmt.Println("Generating thumbnail...")
	os.Chdir(targetDirecotry)

	/*
		-i = Inputfile name
		-vframes 1 = Output one frame
		-an = Disable audio
		-s 400x222 = Output size
		-ss 30 = Grab the frame from 30 seconds into the video
	*/

	cmd := exec.Command(
		"ffmpeg",
		"-i",
		file,
		"-ss",
		"00:00:1.000",
		"-vframes",
		"1",
		"-s",
		"400x222",
		"thumb.jpg",
	)

	cmd.Run()
}
