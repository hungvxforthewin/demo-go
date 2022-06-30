package main

import "fmt"

const MAX int = 3

func main() {
	// Go Pointers in Detail
	// Array of pointers
	a := []int{10, 200, 1000}
	var ptr [MAX] *int
	var i int
	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, a[i] )
	}
	fmt.Printf("\n")
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* assign the address of integer. */
	}
	for  i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *ptr[i] )
	}
	for  i = 0; i < MAX; i++ {
		//fmt.Printf("Address of ptr[%d] = %d\n", i, &ptr[i] ) // Address of pointer
		fmt.Printf("Address of ptr[%d] = %d\n", i, ptr[i] )
		fmt.Printf("Address of a[%d] = %d\n", i, &a[i] )
	}

	// pointer to pointer
	// example1

}