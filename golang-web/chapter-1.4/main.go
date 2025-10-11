package main 

import(
	"fmt"
	"net/http"
	"html/template"

)

type Info struct {
	Affiliation string
	Address string
}

type Person struct {
	Name string
	Gender string 
	Hobbies []string
	Info Info
}

type M map[string]interface{}

func (t Info) GetAffiliationDetailInfo() string {
 return "have 31 divisions"
}


func main() {


	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request){
		var data = M{"name": "Batman"}
		var tmpl = template.Must(template.ParseFiles(
			"views/index.html",
 			"views/_header.html",
 			"views/_message.html",

		))
		var err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	
	 http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
 		var data = M{"name": "Batman"}
 		var tmpl = template.Must(template.ParseFiles(
 					"views/about.html",
 					"views/_header.html",
 					"views/_message.html",
 				))
		 var err = tmpl.ExecuteTemplate(w, "about", data)
		 if err != nil {
 		http.Error(w, err.Error(), http.StatusInternalServerError)
 		}
 		})

	
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			var person = Person{
				Name: "Bruce Wayne",
				Gender: "Male",
				Hobbies: []string{"Reading Books", "Traveling", "Buying things"},
				Info: Info{"Wayne Enterprises", "Gotham City"},
			}

			var tmpl = template.Must(template.ParseFiles("views/template.html"))

			if err := tmpl.Execute(w, person); err != nil {
				 http.Error(w, err.Error(), http.StatusInternalServerError)
			}

		})


	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", nil)

}