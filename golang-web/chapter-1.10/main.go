package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func main() {
	 http.HandleFunc("/", routeIndexGet)
	 http.HandleFunc("/process", routeSubmitPost)
 	 fmt.Println("server started at localhost:9000")
 	 http.ListenAndServe(":9000", nil)

}

func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("view.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)

}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("view.html"))
			
		if err := r.ParseForm(); err != nil {
 			http.Error(w, err.Error(), http.StatusInternalServerError)
 			return
 			}

		var name = r.FormValue("name")
		var message = r.FormValue("message")
		
		var data = map[string]string{"name" : name, "message" : message}
		
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return


	}

	http.Error(w, "", http.StatusBadRequest)
}