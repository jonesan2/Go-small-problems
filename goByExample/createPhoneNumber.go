package kata

import "strconv"

func st(digit uint) string {
	return strconv.Itoa(int(digit))
}

func CreatePhoneNumber(numbers [10]uint) (result string) {
	result += "("
	result += st(numbers[0]) + st(numbers[1]) + st(numbers[2])
	result += ") "
	result += st(numbers[3]) + st(numbers[4]) + st(numbers[5])
	result += "-"
	result += st(numbers[6]) + st(numbers[7]) + st(numbers[8]) + st(numbers[9])
	return
}
