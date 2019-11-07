package main

import (
	"fmt"
	funs "github.com/nattigy/go_tutorial/functions"
)

func main() {
	//closure()
	//structs()
	interfaces()
	//funcValues()
}

func fibo(){
	f := funs.Fibonachi()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func closure(){
	sum := funs.Adder()
	for i := 0; i < 10; i++  {
		fmt.Println(sum(i))
	}
}

func structs(){
	p := funs.Person{
		FirstName: "nathnael",
		LastName: "akale",
		Age: 21,
	}
	fmt.Println(p)
	fmt.Println(p.Greet())
	p.HasBirthday()
	fmt.Println(p)
}

func interfaces(){
	circle := funs.Circles{Radius: 5}
	rectangle := funs.Rectangle{Width: 5, Height: 4}
	fmt.Printf("Area of a circle with radius 5 %v \nArea of a rectangle %v", funs.GetArea(circle), funs.GetArea(rectangle))
}

func funcValues(){
	fmt.Println(funs.FuncValues())
}