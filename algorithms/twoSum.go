package main

import "fmt"

func twoSum(nums *[]int, target int) (int, int) {
	idxMap := make(map[int]int)
	for idx, num := range *nums {
		if idxMap[num] == 0 {
			idxMap[target-num] = idx + 1
		} else {
			return idxMap[num], idx + 1
		}
	}
	return 0, 0
}

func main() {
	var nums = []int{5, 3, 6, 4, 7, 8, 2, 9}
	fmt.Println(twoSum(&nums, 16)) // 5, 8
	fmt.Println(twoSum(&nums, 17)) // 6, 8
	fmt.Println(twoSum(&nums, 5))  // 2, 7
	fmt.Println(twoSum(&nums, 6))  // 4, 7
	nums = []int{-3, -6, 4, 7, 8, 2, 9}
	fmt.Println(twoSum(&nums, -9)) // 1, 2
	fmt.Println(twoSum(&nums, -1)) // 1, 6
}
