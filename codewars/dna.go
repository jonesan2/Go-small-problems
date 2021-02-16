package kata

func DNAStrand(dna string) string {
  dnaMap := map[string]string{"A": "T", "C": "G", "G": "C", "T": "A"}
  var result string
  for _, b := range dna {
    result = result + dnaMap[string(b)] 
  }
  return result
}