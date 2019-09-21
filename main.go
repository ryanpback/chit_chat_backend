package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", welcome)
	log.Fatal(http.ListenAndServe(":3001", mux))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func setUpResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func extractRequestBody(r *http.Request) map[string]interface{} {
	var body = make(map[string]interface{})

	json.NewDecoder(r.Body).Decode(&body)

	return body
}

func welcome(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	setUpResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	request := extractRequestBody(r)

	log.Print(request["data"])

	// data["hello"] = "Hola"
	// json.NewEncoder(w).Encode(data)
}
