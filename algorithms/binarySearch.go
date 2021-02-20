package main

import "fmt"

func binSearch(target int, sortedNums []int) int {
	var idx = len(sortedNums) / 2
	if sortedNums[idx] > target {
		idx = binSearch(target, sortedNums[0:idx])
	} else if sortedNums[idx] < target {
		idx = binSearch(target, sortedNums[idx+1:len(sortedNums)]) + idx + 1
	}
	return idx
}

func main() {
	nums := []int{2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Should be 3:", binSearch(5, nums))
	fmt.Println("Should be 6:", binSearch(8, nums))
	fmt.Println("Should be 5:", binSearch(7, nums))
	fmt.Println("Should be 0:", binSearch(2, nums))
	fmt.Println("Should be 1:", binSearch(3, nums))
}
