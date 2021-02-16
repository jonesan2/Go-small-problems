package main

import "fmt"

func main() {

	var a = "initial" // declare
	fmt.Println(a)

	var b, c int = 1, 2 // declare multiple
	fmt.Println(b, c)

	var d = true // infer type
	fmt.Println(d)

	var e int // zero-valued
	fmt.Println(e)

	f := "apple" // declare and initialize
	fmt.Println(f)
}
