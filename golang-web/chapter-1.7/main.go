package main 

import (
	"fmt"
	"net/http"
	"html/template"
)


type Superhero struct {
	Name string
	Alias string
	Friends []string
}

func (s Superhero) SayHello(from string, message string) string {
	return fmt.Sprintf("$s said: \" $s\"", from, message)
}

 var funcMap = template.FuncMap{
		"unescape": func(s string) template.HTML {
			return template.HTML(s)
		},
		"avg" : func(n ...int) int {
			var total = 0
			for _, each := range n {
				total += each
			}
			return total / len(n)
		},

	 }

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Superhero{
			Name: "Bruce Wayne",
 			Alias: "Batman",
 			Friends: []string{"Superman", "Flash", "Green Lantern"},
		}
		var tmpl = template.Must(template.New("index.html").
 					Funcs(funcMap).
 					ParseFiles("index.html"))

		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}