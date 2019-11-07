package functions

func Fibonachi() func() int {
	num1 := 0
	num2 := 1
	sum := 0
	return func() int {
		num1 = num2
		num2 = sum
		sum = num1 + num2
		return sum
	}
}
