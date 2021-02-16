package kata

import (
  "fmt"
  "strings"
)

func PrinterError(s string) string {
  denom := len(s)
  var numer int
  for _, c := range s {
    if !strings.Contains("abcdefghijklm", string(c)) {
      numer++
    }
  } 
  return fmt.Sprintf("%d/%d", numer, denom)
}