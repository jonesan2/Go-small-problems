// more info on slices: https://blog.golang.org/slices-intro

package main

import "fmt"

func main() {

	// unlike arrays, slices are typed only by the type of their elements,
	//   not the number of elements they contain
	// here is a slice of three zero-valued strings
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// get, set, and len work just like arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// append returns a slice containing one or more new values
	// append requires accepting a return value, because it may
	//     create a new underlying array with a new reference
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy from s into c
	// if c is shorter than s, the copy will have the end cut off
	// if c is longer than s, the end of c will not be overwritten
	c := make([]string, len(s)+1)
	c[len(c)-1] = "test"
	copy(c, s)
	fmt.Println("cpy:", c)
	fmt.Println("cpy:", s)

	// slice operator: slice[low:high]
	// slice will begin at low and go up to, but excluding, high
	l := s[2:5]
	fmt.Println("sl1:", l) // prints [c d e]

	l = s[:5]
	fmt.Println("sl2:", l) // prints [a b c d e]

	l = s[2:]
	fmt.Println("sl3:", l) // prints [c d e f]

	// declare and initialize in one line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// in multiD slices, the length of the inner slices can vary (unlike arrays)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
