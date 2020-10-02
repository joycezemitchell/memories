package gd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ConvertToHls file
func ConvertToHls(file string, targetDirecotry string) {
	// ffmpeg  -i VID_20191227_204231.mp4 -vf scale=w=1920:h=1080:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -c:v h264 -profile:v main -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod -b:v 5000k -maxrate 5350k -bufsize 7500k -b:a 192k -hls_segment_filename 1080p_%d.ts 1080p.m3u8

	// Ref: https://docs.peer5.com/guides/production-ready-hls-vod/
	os.Chdir(targetDirecotry)

	fmt.Println("Get video rotaion")

	videoRotation := getVideoRotation(file, targetDirecotry)

	if videoRotation == "Portrait" {
		// Convert video to landscape first
		fmt.Println("Converting video to landcape")
		convertToLandscape(file, targetDirecotry)
		file = "outputfile.mp4"
		GenerateThumbnails(file, targetDirecotry)
	}

	fmt.Println("Generating HLS files...")

	// Needed to add this for MOV, somehow, w=1920:h=1080 is not working.
	// TODO: trace the actual problem why ffmpeg is throwing an error when using MOV

	/*scale := "scale=w=1920:h=1080:force_original_aspect_ratio=decrease"

	file = strings.ToLower(file)
	i := strings.Index(file, ".mov")
	if i > -1 {
		scale = "scale=w=1280:h=720:force_original_aspect_ratio=decrease"
	}*/

	scale := "scale=1920:-2"

	cmd := exec.Command(
		"ffmpeg",
		"-i", file,
		"-vf", scale,
		"-c:a", "aac",
		"-ar", "48000",
		"-c:v", "h264",
		"-profile:v", "main",
		"-crf", "20",
		"-sc_threshold", "0",
		"-g", "48",
		"-keyint_min", "48",
		"-hls_time", "4",
		"-hls_playlist_type", "vod",
		"-b:v", "5000k",
		"-maxrate", "5350k", "-bufsize", "7500k", "-b:a",
		"192k", "-hls_segment_filename", "1080p_%d.ts", "1080p.m3u8",
	)

	cmd.Run()
}

func getVideoRotation(file string, targetDirecotry string) string {

	os.Chdir(targetDirecotry)

	cmdName := `exiftool ` + file
	cmdArgs := strings.Fields(cmdName)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	out, _ := cmd.CombinedOutput()
	s := string(out)
	b := strings.Split(s, "\n")
	var i int

	for _, n := range b {
		i = strings.Index(n, "Rotation")
		if i > -1 {
			fmt.Println(n)
			i = strings.Index(n, "90")
			if i > -1 {
				fmt.Println("Portrait")
				return "Portrait"

			} else {
				fmt.Println("Landscape")
				return "Landscape"

			}
		}
	}

	return ""

}

func convertToLandscape(file string, targetDirecotry string) {

	os.Chdir(targetDirecotry)

	param := `[0:v]scale=ih*16/9:-1,boxblur=luma_radius=min(h\,w)/20:luma_power=1:chroma_radius=min(cw\,ch)/20:chroma_power=1[bg];[bg][0:v]overlay=(W-w)/2:(H-h)/2,crop=h=iw*9/16`

	cmdName := "ffmpeg -y -i " + file + " -filter_complex " + param + " outputfile.mp4"
	cmdArgs := strings.Fields(cmdName)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	cmd.Run()

}
