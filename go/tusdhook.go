package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
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
	ioutil.WriteFile("test101.txt", body, 0644)

	// - parse json - parse filename

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

	if header == "post-finish" {

		os.Chdir(tusStorage)
		cmd := exec.Command("cp", tusFilename, tusTempDirectory+"/tmphls/"+originalFilename)
		cmd.Stdout = os.Stdout
		cmd.Run()

	}

	// - check if directory exist in do space based on filename
	// s3cmd ls s3://allyapps | grep 134

	// out, _ := exec.Command("s3cmd", "ls", "s3://allyapps").Output()
	// fmt.Println(string(out))

}
