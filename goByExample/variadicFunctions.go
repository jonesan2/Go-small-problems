// variadic functions can be called with any number of trailing arguments

package main

import "fmt"

// this function accepts an arbitrary number of int arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// if you already have multiple args in a slice, apply them to a variadic
	//    function using func(slice...) like this:
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
