package kata

import "strings"

func duplicate_count(s1 string) (count int) {
  charMap := make(map[string]int)  
  for _, c := range s1 {
    charMap[strings.ToLower(string(c))] += 1
  }
  for _, v := range charMap {
    if v > 1 {
      count++
    }
  }
  return
}