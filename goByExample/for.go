package main

import "fmt"

func main() {

	// for is Go's only looping construct

	// most basic type
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic type
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// loop repeatedly until break or return
	for {
		fmt.Println("loop")
		break
	}

	// continue to next iteration
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
