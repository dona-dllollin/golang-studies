package main 

import (
	"log"
	"net/http"
	"strings"
	"time"
	"io"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)
	go func() {
		// do the process here
		// simulate a long-time request by putting 10 seconds sleep

		 body, err := io.ReadAll(r.Body)
			_ = err
			_ = body

		time.Sleep(10 * time.Second)

		done <- true
	}()
	
	select {
		case <-r.Context().Done():
			if err := r.Context().Err(); err != nil {
				if strings.Contains(strings.ToLower(err.Error()), "canceled") {
					log.Println("Request is canceled by the client")
				} else {
					log.Println("Unknown error:", err.Error())
				}
			}
			
		case <-done:
			log.Println("Process completed")
}
}


func main() {
 http.HandleFunc("/", handleIndex)
 http.ListenAndServe(":8080", nil)
}