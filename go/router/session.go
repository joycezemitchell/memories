package router

import (
	"net/http"
	"encoding/json"
)

// Error message for JSON API Request
type errorMessage struct {
	ResponseMessage string
	ErrorMessage    string
}


// Session
func Session() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		switch r.Method {
		case http.MethodPost:			
		case http.MethodGet:
		default:
		
		}
		errormessage := errorMessage{}
		errormessage.ResponseMessage = ""
		errormessage.ErrorMessage = ""
	
		json.NewEncoder(w).Encode(errormessage)


	})
}