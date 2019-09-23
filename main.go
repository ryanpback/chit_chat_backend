package main

import (
	"chitChat/requesthandler"
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
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func welcome(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	setUpResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	var data map[string]interface{}
	if err := requesthandler.Decode(r, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
