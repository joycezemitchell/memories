package gd

import (
	"fmt"
	"strings"
	"os"
)	

// GetDate 
func GetDate(info os.FileInfo) string {

	st := fmt.Sprintf("%v", info.ModTime())
	s := strings.Split(st, " ")
	date := s[0]
	
	return date
}
