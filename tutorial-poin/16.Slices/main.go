package main

import "fmt"

func main() {
	// Defining a slice
	var numbers []int /* a slice of unspecified size */
	/* numbers == []int{0,0,0,0,0}*/
	numbers = make([]int, 5, 5) /* a slice of length 5 and capacity 5*/
	printSlice(numbers)

	// Nil slice

	var mySilce []int
	printSlice(mySilce)

	if(mySilce == nil){
		fmt.Printf("slice is nil")
	}
}

// len(): thực tế, cap(): có thể chứa
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}