package main

import "fmt"

func main() {
	fmt.Println("Constants")
	/*
		Integer Literals
	*/
	const a1 int = 212
	const a2 uint64 = 2
	/*
		Floating-point Literals
	*/
	const a3 float64 = 3.14159
	/*
		Escape Sequence
	*/
	fmt.Println("Hello\\World!")
	fmt.Println("Hello\"World!")
	fmt.Println("Hello\\'World!")
	fmt.Println("Hello\\?World!")
	fmt.Println("Hello\aWorld!")
	fmt.Println("Hello\bWorld!")
	fmt.Println("Hello\fWorld!")
	fmt.Println("Hello\nWorld!")
	fmt.Println("Hello\rWorld!")
	fmt.Println("Hello\tWorld!")
	fmt.Println("Hello\vWorld!")

}
