package main 

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func main() {

	mux := new(CustomMux)

	mux.HandleFunc("/student", ActionStudent)

	
	mux.RegisterMiddleware(MiddlewareAuth)
	mux.RegisterMiddleware(MiddlewareAllowOnlyGet)

	server := new(http.Server)
	server.Addr = ":9000"
	server.Handler = mux

	fmt.Println("Starting server on port 9000...")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, SelectAllStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}