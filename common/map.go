package kata

func Maps(x []int) (result []int) {
  for _, n := range x {
    result = append(result, n * 2)
  } 
  return result
}
