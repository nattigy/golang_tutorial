package questions

import (
	"fmt"
	"math/rand"
	"time"
)

func Guess()  {
	fmt.Println("Please guess a number ...")
	rand.Seed(time.Now().UnixNano())
	secretNum := rand.Intn(20)
	var entries [6]int
	for i := 0; i <= 5; i++ {
		var input int
		fmt.Scanln(&input)
		entries[i] = input
		if input < secretNum {
			fmt.Println("Too low...")
		} else if input > secretNum {
			fmt.Println("Too high...")
		} else if input == secretNum {
			fmt.Println("Congrats you got the secret number")
		}
		if check(entries, input){
			//fmt.Println(entries)
		}
		if i == 5 && input != secretNum || input == secretNum {
			fmt.Println("Game over... \n 1. Play again \n 2. Exit")
			var x int
			fmt.Scanln(&x)
			if x == 1 {
				fmt.Println("Please guess a number ...")
				i = 0
			} else {
				break
			}
		}

	}
}

func check(a [6]int, input int) bool{
	for i := 0; i < 6; i++{
		if a[i] == input{
			fmt.Println(a)
			return true
		}
	}
	return false
}
