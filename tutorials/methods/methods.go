package main

import "fmt"

type rect struct {
	width, height int
}

//*rect is a pointer receiver type
func (r *rect) area() int {
	return r.width * r.height
}

//rect is a value receiver type
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	//This is different. Functions with
	//the rect receiving type are automatically
	//available in the variable..similar to a class.
	//It's as if a class is automatically created.
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
