package main

import "fmt"

func spiralMatrix(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	var j, i int
	var result []int
	var iSteps, jSteps = m - 1, n
	var count, totalCount int
	for {
		for count = 1; count <= jSteps; count++ {
			result = append(result, matrix[i][j])
			j++
			totalCount++
		}
		if totalCount >= m*n {
			break
		}
		jSteps--
		j--
		i++
		for count = 1; count <= iSteps; count++ {
			result = append(result, matrix[i][j])
			i++
			totalCount++
		}
		if totalCount >= m*n {
			break
		}
		iSteps--
		i--
		j--
		for count = 1; count <= jSteps; count++ {
			result = append(result, matrix[i][j])
			j--
			totalCount++
		}
		if totalCount >= m*n {
			break
		}
		jSteps--
		j++
		i--
		for count = 1; count <= iSteps; count++ {
			result = append(result, matrix[i][j])
			i--
			totalCount++
		}
		if totalCount >= m*n {
			break
		}
		iSteps--
		i++
		j++
	}
	return result
}

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	fmt.Println(spiralMatrix((mat)))
}
