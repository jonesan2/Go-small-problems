package kata

import (
  "fmt"
  "strings"
  "strconv"
)

func HighAndLow(in string) string {
  nums := strings.Split(in, " ") 
  low, _ := strconv.Atoi(nums[0])
  high := low
  for _, n := range nums {
    n, _ := strconv.Atoi(n)
    if n < low {
      low = n
    } else if n > high {
      high = n
    }
  }
  
  return fmt.Sprintf("%d %d", high, low)
}