package questions

import "fmt"

func Diamond(n int)  {
	space := n - 1
	for i := 0; i < n; i++ {
		for j := 0; j < space; j++ {
			fmt.Print(" ")
		}

		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}

		fmt.Println()
		space--
	}

	space = 0
	for i := n; i > 0; i-- {
		for j := 0; j < space; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}

		fmt.Println()
		space++
	}
}
