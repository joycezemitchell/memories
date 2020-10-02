package gd

import (
	"strings"
	"time"
)

// GenerateDirectoryNameFromFilename file
func GenerateDirectoryNameFromFilename(file string) string {

	fileSlice := strings.Split(file, ".")
	dirName := strings.Join(fileSlice, "")
	dt := time.Now()

	dirName = dt.Format("01022006") + "-" + dirName
	return dirName
}
