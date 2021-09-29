package main

import "fmt"

type Animal interface {
	Speak() string
}

//implement Dog
type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

//implement cat
type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow"
}

func main() {

	animals := []Animal{Dog{}, Cat{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

}
