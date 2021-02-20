//__________/\\\_______________________________________/\\\\\\\\\\\\\\\_
//_________/\\\\\______________________________________\/\\\///////////__
//________/\\\/\\\___________________/\\\_______________\/\\\_____________
// ______/\\\/\/\\\_____/\\\\\\\\\\\_\///___/\\\\\\\\\\\_\/\\\\\\\\\\\_____
//  ____/\\\/__\/\\\____\///////////___/\\\_\///////\\\/__\/\\\///////______
//   __/\\\\\\\\\\\\\\\\_______________\/\\\______/\\\/____\/\\\_____________
//    _\///////////\\\//________________\/\\\____/\\\/______\/\\\_____________
//     ___________\/\\\__________________\/\\\__/\\\\\\\\\\\_\/\\\\\\\\\\\\\\\_
//      ___________\///___________________\///__\///////////__\///////////////__

package main

import "fmt"

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}
	return sum
}

func maxValue(n1 int, n2 int, n3 int) int {
	if n1 > n2 && n1 > n3 {
		return n1
	} else if n2 > n3 {
		return n2
	} else {
		return n3
	}
}

func maxLeftEdge(nums []int) int {
	curMax := nums[0]
	for idx := len(nums) - 1; idx >= 0; idx-- {
		curSum := sum(nums[idx:len(nums)])
		if curMax < curSum {
			curMax = curSum
		}
	}
	return curMax
}

func maxRightEdge(nums []int) int {
	curMax := nums[len(nums)-1]
	for idx := 0; idx < len(nums); idx++ {
		curSum := sum(nums[0 : idx+1])
		if curMax < curSum {
			curMax = curSum
		}
	}
	return curMax
}

func maxSubSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	length := len(nums)
	mid := length / 2
	leftArray := nums[0:mid]
	rightArray := nums[mid:length]
	leftMax := maxSubSum(leftArray)
	rightMax := maxSubSum(rightArray)
	centerMax := maxLeftEdge(leftArray) + maxRightEdge(rightArray)
	return maxValue(leftMax, rightMax, centerMax)
}

func main() {
	test1 := []int{-2, -5, 6, -2, -3, 1, 5, -6}
	fmt.Println(maxSubSum(test1))
}
