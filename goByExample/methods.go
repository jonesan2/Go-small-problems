// Go supports methods defined on struct types

package main

import "fmt"

type rect struct {
	width, height int
}

// this method has a receiver type of *rect
// the space after the func keyword is crucial
//    that space is what indicates that this is a method on the rect type
//    without the space, it would be an anonymous function with a rect param
func (r *rect) area() int {
	return r.width * r.height
}

// methods can be defined for either pointer or value receiver types
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Go automatically handles conversion between values and pointers
	//     for method calls. You may want to use a pointer receiver type
	//     to avoid copying on method calls or to allow the method to
	//     mutate the receiving struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
