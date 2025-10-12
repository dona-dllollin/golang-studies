package main 

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/", ActionIndex)
	http.HandleFunc("/encode", ActionIndexEncode)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)

}

func ActionIndex(w http.ResponseWriter, r *http.Request){
	data := []struct{
		Name string
		Age int
	}{
		{"Richard Grayson", 24},
		{"Jason Todd", 23},
		{"Tim Drake", 22},
		{"Damian Wayne", 21},
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)

}
func ActionIndexEncode(w http.ResponseWriter, r *http.Request){
	data := []struct{
		Name string
		Age int
	}{
		{"Richard Grayson", 24},
		{"Jason Todd", 23},
		{"Tim Drake", 22},
		{"Damian Wayne", 21},
	}

	w.Header().Set("Content-Type", "application/json")
	if err:= json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
 		return

	}

}