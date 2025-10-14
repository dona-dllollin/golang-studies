package main 

import "net/http"

const USERNAME = "admin"
const PASSWORD = "admin"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth() 
	
	if !ok {
		w.Write([]byte(`something went wrong`))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte(`invalid credentials`))
		return false
	}

	return true
}

func AllowOnlyGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte(`only GET method is allowed`))
		return false
	}
	return true
}

