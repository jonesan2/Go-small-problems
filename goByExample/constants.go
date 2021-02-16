package main

import (
	"fmt"
	"math"
)

const s string = "constant" // declare constant

func main() {

	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n // arbitrary precision
	fmt.Println(d)

	fmt.Println(int64(d)) // numeric constant has no type until given one

	fmt.Println(math.Sin(n)) // using a constant number in a context that requires
	// a certain type will give it that type. e.g. math.Sin expects a float64
}
