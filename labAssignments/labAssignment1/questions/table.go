package questions

import "fmt"

func Table(){
	for i := 1; i <= 12; i++ {
		for j := 1; j <= 12; j++ {
			fmt.Print(" ", i * j, " ")
		}
		fmt.Println()
	}
}
