package kata

import (
	"strconv"
	"strings"
)

func NbDig(n int, d int) (count int) {
	var squares []string
	for i := 0; i <= n; i++ {
		squares = append(squares, strconv.Itoa(i*i))
	}
	squaresString := strings.Join(squares, "")
	for _, v := range squaresString {
		if string(v) == strconv.Itoa(d) {
			count++
		}
	}
	return count
}

/*
func NbDig(n int, d int) (count int) {
	for i := 0; i <= n; i++ {
		out += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}
}
*/
