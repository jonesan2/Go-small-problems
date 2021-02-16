package main

import "fmt"

func main() {

	// array will hold exactly 5 ints
	// zero-valued by default
	var a [5]int
	fmt.Println("emp:", a)

	// set and get
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// built-in function len
	fmt.Println("len:", len(a))

	// declare and initialize in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// compose types to build multi-dimensional data structures
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	// arrays appear in the form [v1 v2 v3] when printed with fmt.Println
	fmt.Println("2d: ", twoD)
}
