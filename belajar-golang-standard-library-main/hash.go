package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func main() {
	// var text = "this is secret"
	// var sha = sha1.New()
	// sha.Write([]byte(text))
	// var encrypted = sha.Sum(nil)
	// var encryptedString = fmt.Sprintf("%x", encrypted)
	// fmt.Println(encryptedString)
	// f4ebfd7a42d9a43a536e2bed9ee4974abf8f8dc8

	var text = "this is secret"
	fmt.Printf("original : %s\n\n", text)
	var hashed1, _ = doHashUsingSalt(text)
	fmt.Printf("hashed 1 : %s\n\n", hashed1)
	// 929fd8b1e58afca1ebbe30beac3b84e63882ee1a
	var hashed2, _ = doHashUsingSalt(text)
	fmt.Printf("hashed 2 : %s\n\n", hashed2)
	// cda603d95286f0aece4b3e1749abe7128a4eed78
	var hashed3, _ = doHashUsingSalt(text)
	fmt.Printf("hashed 3 : %s\n\n", hashed3)
	// 9e2b514bca911cb76f7630da50a99d4f4bb200b4
	// _, _, _ = salt1, salt2, salt3

}

func doHashUsingSalt(text string) (string, string) {
	salt := fmt.Sprintf("%d", time.Now().UnixNano())

	saltText := fmt.Sprintf("text : '%s', salt: %s", text, salt)

	var sha = sha1.New()

	sha.Write([]byte(saltText))

	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}
