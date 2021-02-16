package main

import "fmt"

func plus(a int, b int) int {

	// explicit returns are required
	return a + b
}

// with multiple consecutive parameters of the same type, you may omit the
//    type name until after the final like-typed parameter
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
