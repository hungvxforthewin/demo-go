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
}