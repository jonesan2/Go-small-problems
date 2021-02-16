package main

import (
	"fmt"
	"sort"
)

// Sort methods are specific to the bulletin type
// Sorting is in-place

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:    ", ints)

	// to check if a slice is already sorted:
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
