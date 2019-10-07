package main

import (
	"fmt"
)

func fibonachi() func() int {
	num1 := 0
	num2 := 1
	sum := 0
	return func () int {
		num1 = num2
		num2 = sum
		sum = num1 + num2
		return sum
	}
}

func maifn() {
	f := fibonachi()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}