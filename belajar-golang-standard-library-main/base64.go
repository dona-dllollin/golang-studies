package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Eko Kurniawan Khannedy"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}

	var data = "https://kalipare.com/"
	var encodedString = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString)
	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println(decodedString)
}
