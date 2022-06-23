package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", helloWordCallerHandler)
	http.ListenAndServe(":8080", serveMux)
}

func helloWordCallerHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Method: %s Path: %s", r.Method, r.URL.Path)

	var result map[string]string

	type Body struct {
		Result map[string]string `json:"result"`
		Err    error             `json:"err"`
	}

	response, err := http.Get("http://simple-api/")
	if err == nil {
		json.NewDecoder(response.Body).Decode(&result)
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(Body{result, err})
}
