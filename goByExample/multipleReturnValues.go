package main

import "fmt"

// this function returns 2 ints
func vals() (int, int) {
	return 3, 7
}

func main() {
	// multiple assignment
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// blank identifier, used because we only want a subset of the return values
	_, c := vals()
	fmt.Println(c)
}
