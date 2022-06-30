package main

import "fmt"

func main() {
	// Address variable
	var a int = 10
	fmt.Printf("Address of a variable: %x\n", &a)
	// Pointer
	var ip *int
	ip = &a

	fmt.Printf("Address of a variable: %x\n", &a)

	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip)

	// nil Pointer
	var ptr *int
	fmt.Printf("The value of ptr is : %x\n", ptr)
	if ptr == nil {
		fmt.Printf("The value of ptr is nil")
	}
}
