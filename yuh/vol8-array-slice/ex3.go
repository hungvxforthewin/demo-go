package main

import (
	"fmt"
)

func main() {
	a := make([]int, 0)
	a = append(a, 0)
	a = append(a, 1)
	a = append(a, 3)
	a = append(a, 4)
	a = append(a, []int{7,7,7}...)

	fmt.Printf("%v %v, %v\n", a, len(a), cap(a))
}