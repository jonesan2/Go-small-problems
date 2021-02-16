package kata

func Tribonacci(signature [3]float64, n int) []float64 {
	result := make([]float64, 0)
	if n > 0 {
		result = append(result, signature[0])
		if n > 1 {
			result = append(result, signature[1])
			if n > 2 {
				result = append(result, signature[2])
				if n > 3 {
					for i := 3; i < n; i++ {
						nextVal := result[i-1] + result[i-2] + result[i-3]
						result = append(result, nextVal)
					}
				}
			}
		}
	}
	return result
}
