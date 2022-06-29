package main

import "fmt"

/* global variable declaration */
var g int
var g1 int = 7

func main() {
	fmt.Println("Welcome to the OpenStack Project - Scope Rules")
	// Local Variables and Global Variables, parameter
	// Scope Rules

	// Local Variables
	/* local variable declaration */
	var a, b, c int
	var g1 int = 10
	/* actual initialization */
	a = 10
	b = 20
	c = a + b
	g = a + b + c

	fmt.Printf("value of a = %d, b = %d and c = %d\n", a, b, c)
	fmt.Printf("value of a = %d, b = %d and g = %d\n", a, b, g)
	fmt.Printf("value of g = %d\n", g1)
	d := sum( a, b)
	fmt.Printf("value of d in main() = %d\n",  d)
}

/* function to add two integers */
func sum(a, b int) int {
	fmt.Printf("value of a in sum() = %d\n",  a)
	fmt.Printf("value of b in sum() = %d\n",  b)

	return a + b
}