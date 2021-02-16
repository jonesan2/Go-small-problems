// example of custom sorting in Go

package main

import (
	"fmt"
	"sort"
)

// first you need a custom type
type byLength []string

// then you need to define Len, Swap, and Less methods on it

// Swap and Len will usually be the same across types,
//     but you still have to define them (that's annoying)
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less holds the actual sorting logic
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	// Defining Len, Swap, and Less on your custom type allows you to call the
	//      sort package's generic Sort function on it
	// In this case we convert a slice of strings to a slice of our custom type
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
