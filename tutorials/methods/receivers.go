package main

import "fmt"

type person struct {
	name string
	age  int
}

//the (p person) makes this function a
//receiver function. It can receive a person.
func (p person) print() {
	fmt.Printf("%s is of %d years \n", p.name, p.age)
}

func main() {
	alex := person{
		name: "Alex",
		age:  18,
	}

	//the receiver function shows up as part
	//of the variable.
	alex.print()
}
