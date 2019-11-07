package functions

import "fmt"

func pointers() {
	i, j := 1, 2
	p := &i
	*p = j
	fmt.Println(i)
	var t = 0
	fmt.Println(*&t)
	fmt.Println(My{"hello", "world"})

	/* Arrays */
	var a [2]string
	/* Or */
	var myArray = [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArray, a)

	/* Or */
	var arr = myArray[1:2]
	fmt.Println(arr)

	/* More on Struct*/
	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{0, false},
	}
	fmt.Println(s)
}

type My struct {
	X, Y string
}
