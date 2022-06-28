package main

import "fmt"

func main(){
	// Static Type Declaration in Go
	var x float64
	x = 20.0
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)
	// Dynamic Type Declaration
	y := 42
	fmt.Println(y)
	fmt.Println("y is of type %T\n", y)
	// Mixed Variable Declaration in Go
	var a, b, c = 	0, 0, "hello world"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
	// The lvalues and the rvalues in Go
	/*
		Variable = value
	*/
}