package main

import (
	"fmt"
	"strings"
)

type MyFloat int

type Person struct {
	firstName string
	lastName string
	gender string
	age int
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(1)
	var n = "hehe"
	var person1 Person
	person1.firstName = "nathnael"
	person1.lastName = "akale"
	person1.gender = "m"
	person1.age = 21
	fmt.Println(person1)
	fmt.Println(strings.ToUpper(n))
	fmt.Println(f.Abs())
}
