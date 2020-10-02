package gd

import (
	"fmt"
	"os"
	"os/exec"
)

// Upload file
func Upload(file string, targetDirecotry string, newDirectory string) {
	fmt.Println("Uploading", file, "...")
	os.Chdir(targetDirecotry)

	// TUSD does not seem to be working. I am getting disconnected when trying to run this command.
	// DEBUG=1 tusc --host http://memories.allyapps.com directory_filename
	/*cmd := exec.Command(
		"tusc",
		"--host",
		"http://memories.allyapps.com",
		file,
	)*/

	// s3cmd put text2.txt s3://allyapps/ --acl-public
	cmd := exec.Command(
		"s3cmd",
		"put",
		file,
		"s3://allyapps/"+newDirectory+"/",
		"--acl-public",
	)

	cmd.Run()
}
