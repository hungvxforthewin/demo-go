package main

import (
	"fmt"
)

func main() {
	// pointer array
	//a := [3]int{1, 2, 3}
	//b := &a
	//b := a // copy
	//b[0] = 10
	//fmt.Println(a)
	//fmt.Println(b)

	// slice
	a := []int{1, 2, 3}
	b := a
	fmt.Println(a)
	fmt.Println(len(b))
	fmt.Println(cap(b))

}
