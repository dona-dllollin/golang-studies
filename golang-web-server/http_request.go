package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type studentRes struct {
	ID    string
	Name  string
	Grade int
}

func fetchUsers() ([]studentRes, error) {
	var err error
	var client = &http.Client{}
	var data []studentRes

	request, err := http.NewRequest("GET", "http://localhost:8080/users", nil)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		return nil, err
	}
	return data, nil

}

var baseURL = "http://localhost:8080"

func fetchUser(ID string) (studentRes, error) {
	var err error
	var client = &http.Client{}
	var data studentRes

	var param = url.Values{}

	param.Set("id", ID)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/user", payload)

	if err != nil {
		return data, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		// baca body supaya tahu kenapa error
		body, _ := io.ReadAll(response.Body)
		return data, fmt.Errorf("server error: %s", string(body))
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	var err error
	users, err := fetchUsers()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	for _, each := range users {
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	}

	user1, err := fetchUser("E001")
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade)

}
