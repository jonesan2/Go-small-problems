package kata

import (
  "strings"
)

func reverse(subs string) (reversed string) {
  for _, c := range subs {
    reversed = string(c) + reversed
  }
  return
}
  
func VertMirror(s string) string {
  lines := strings.Split(s, "\n")  
  var mappedLines []string = make([]string, len(lines[0]))
  for i, line := range lines {
    mappedLines[i] = reverse(line) 
  }
  return strings.Join(mappedLines, "\n") 
}

func HorMirror(s string) string {
  lines := strings.Split(s, "\n")  
  result := lines[0]
  for i, line := range lines {
    if i == 0 {
      continue
    }
    result = line + "\n" + result 
  }
  return result
}

type FParam func(string) string

func Oper(f FParam, x string) string {
  return f(x)
}