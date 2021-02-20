package main

import "fmt"

func spiralMatrix(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	var j, i int
	var result []int
	var iSteps, jSteps = m - 1, n
	var count, totalCount int
	var steps, pointer *int

	getLine := func(directions [2]bool) {
		var across = directions[0]
		var positive = directions[1]
		if across {
			steps = &jSteps
			pointer = &j
		} else {
			steps = &iSteps
			pointer = &i
		}

		for count = 1; count <= *steps; count++ {
			result = append(result, matrix[i][j])
			if positive {
				*pointer++
			} else {
				*pointer--
			}
			totalCount++
		}
		*steps--

		if positive {
			j--
			if across {
				i++
			} else {
				i--
			}
		} else {
			j++
			if across {
				i--
			} else {
				i++
			}
		}
	}

	var directions = [4][2]bool{{true, true}, {false, true}, {true, false}, {false, false}}

	var k int
	for {
		getLine(directions[k%4])
		if totalCount >= m*n {
			break
		}
		k++
	}
	return result
}

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	fmt.Println(spiralMatrix((mat)))
}
