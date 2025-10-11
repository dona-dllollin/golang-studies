package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Eko"}
	SayHello(person)

	animal := Animal{Name: "Kucing"}
	SayHello(animal)

	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)
	secret = 12.4
	fmt.Println(secret)

	var number = secret.(int) * 10
	fmt.Println(secret, "multiplied by 10 is :", number)

	type person2 struct {
		name string
		age  int
	}

	var secret2 interface{} = &person2{name: "wick", age: 27}
	var name = secret2.(*person2).name
	fmt.Println(name)

}
