package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {

	var jsonString = `{"Name" : "john wick", "age": 27}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var data1 map[string]interface{}

	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data.FullName)
	fmt.Println("age :", data.Age)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age :", data1["age"])

	var data2 interface{}

	json.Unmarshal(jsonData, &data2)
	fmt.Printf("%+v\n", data2)

	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}

	var jsonData2, error = json.Marshal(object)

	if error != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString2 = string(jsonData2)
	fmt.Println(jsonString2)
}
