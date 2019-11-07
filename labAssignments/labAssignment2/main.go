package main

import "fmt"

func main(){
	combineList()
	f := Fibonachi()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func combineList(){
	array1 := []string{"1","2","3"}
	array2 := []string{"a","b","c"}
	newarr := [6]string{}
	index := 0

	for i := 0; i < 2 * len(array1); i++ {
		if i % 2 == 0 {
			newarr[i] = array1[index]
		} else {
			newarr[i] = array2[index]
			index++
		}
	}

	fmt.Println(newarr)
}

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


