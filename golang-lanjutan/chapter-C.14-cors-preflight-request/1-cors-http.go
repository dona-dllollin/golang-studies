package main

import (
	"log"
	"net/http"
)
func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://www.google.com")
		w.Header().Set("Access-Control-Allow-Methods", "options, POST, GET, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

		if r.Method == "OPTIONS" {
			w.Write([]byte("Allowed"))
			return
		}

		w.Write([]byte("Hello World"))

	})

	log.Println("Starting server at :9000")
	http.ListenAndServe(":9000", nil)
}