package kata

import "strings"

func DuplicateEncode(word string) (result string) {
	charMap := make(map[string]int)
	for _, c := range word {
		charMap[strings.ToLower(string(c))] += 1
	}
	for _, c := range word {
		if charMap[strings.ToLower(string(c))] > 1 {
			result += ")"
		} else {
			result += "("
		}
	}
	return
}
