package primitives

func Abs(number int) int {
	if number < 0 {
		return number * (-1)
	}
	return number
}
