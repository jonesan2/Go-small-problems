package main

import (
	"fmt"
)

/*
func Solution(word string) string {
	if len(word) == 0 {
		return ""
	}
	reversed := make([]string, len(word))
	for i := 0; i <= len(word)/2; i++ {
		reversed[i], reversed[len(word)-1-i] = string(word[len(word)-1-i]), string(word[i])
	}
	return strings.Join(reversed, "")
}
*/

func Solution(word string) (reversed string) {
	for _, c := range word {
		reversed = string(c) + reversed
	}
	return reversed
}

func main() {
	fmt.Println(Solution("wacky"))
}
