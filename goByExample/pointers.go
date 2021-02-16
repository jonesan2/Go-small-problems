package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// the *iptr code DEREFERENCES the pointer from its memory address to
//     the current value at that address
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// the &i syntax gives the memory address of i, i.e. a pointer to i
	fmt.Println("pointer:", &i)

}
