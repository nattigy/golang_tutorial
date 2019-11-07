package functions

import (
	"fmt"
	"math"
)

var python string

func add(x, y int) (massage string, sum int) {
	sum = x + y
	massage = "the sum is"
	return
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func forr(num int) {
	my := num
	fmt.Println("counting")
	defer fmt.Println("end")
	for i := 0; i <= num; i++ {
		defer fmt.Println(my)
		my -= 1
	}
	defer fmt.Println("begin")
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		z -= (z*z - x) / (2 * z)
		if z*z == x {
			return z
		}
		fmt.Println(z)
	}
}

func hello() {
	forr(10)
	fmt.Println(pow(3, 2, 5))
	fmt.Println(Sqrt(10000))
	fmt.Printf("my %g \n", math.Sqrt(64))
	fmt.Println(add(1, 2))
	java := "java"
	fmt.Println(python, java)
	fmt.Printf("Type: %T Value: %q\n", python, python)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
