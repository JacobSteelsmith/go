package main

import (
	"fmt"
	"math"
)

//interface
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

//implement the methods in the interface
//for the rect type
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

//implement the methods in the interface
//for the circle type
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//if a variable (g) has an interface type,
//we can call all methods implemented.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {

	//the variables of the typ
	r := rect{width: 3, height: 4}
	//c := circle{radius: 5}

	//we can call r.area because our area
	//function has a receiver type of rect
	fmt.Printf("r.area is %f\n", r.area())
	fmt.Println("measure(r):")
	measure(r)

}
