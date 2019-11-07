package functions

import (
	"fmt"
	"strings"
)

type MyFloat int

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func methonds() {
	f := MyFloat(1)
	var n = "hehe"
	fmt.Println(strings.ToUpper(n))
	fmt.Println(f.Abs())
}
