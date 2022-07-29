package main

import (
	"fmt"
)

func main() {
	a := []int {1, 2, 3, 4, 5}
	b := a[:4]
	c := a[1:]
	d := append(a[:2],a[3:]...) // thay doi gia tri a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a)
}
