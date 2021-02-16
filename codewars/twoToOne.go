package kata

func TwoToOne(s1 string, s2 string) (result string) {
  lettersMap := make(map[rune]bool)
  for _, c := range s1 + s2 {
    lettersMap[c] = true 
  } 
  for _, c := range "abcdefghijklmnopqrstuvwxyz" {
    if lettersMap[c] {
      result += string(c)
    }
  }
  return
}