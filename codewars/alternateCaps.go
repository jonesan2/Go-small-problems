package kata

import "strings"

func Capitalize(st string) []string {
  var str1, str2 []string
  for i, c := range st {
    if i%2 == 0 {
      str1 = append(str1, strings.ToUpper(string(c))) 
      str2 = append(str2, strings.ToLower(string(c))) 
    } else {
      str2 = append(str2, strings.ToUpper(string(c))) 
      str1 = append(str1, strings.ToLower(string(c))) 
    }
  }
  
  return []string{strings.Join(str1, ""), strings.Join(str2, "")}
}