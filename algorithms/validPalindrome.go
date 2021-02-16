package main

import (
	"fmt"
	s "strings"
)

func movePointers(idxFront int, idxBack int, str string) (int, int) {
	for (str[idxFront] < 65 || str[idxFront] > 90) && idxFront < idxBack {
		idxFront++
	}
	for (str[idxBack] < 65 || str[idxBack] > 90) && idxFront < idxBack {
		idxBack--
	}
	return idxFront, idxBack
}

func validPalindrome(str string) bool {
	if str == "" {
		return true
	}

	str = s.ToUpper(str)
	var idxFront, idxBack = 0, len(str) - 1
	for idxFront <= idxBack {
		idxFront, idxBack = movePointers(idxFront, idxBack, str)
		if str[idxFront] != str[idxBack] {
			return false
		}
		idxFront++
		idxBack--
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("-a: ..A&%# "))
	fmt.Println(validPalindrome("-: ..&%# "))
	fmt.Println(validPalindrome("racecar"))
	fmt.Println(validPalindrome("A man, a plan, a canal, Panama."))
	fmt.Println(validPalindrome("A man, a piano, a canal, Panama."))
	fmt.Println(validPalindrome("race a car"))
}
