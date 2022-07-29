package main

import (
	"fmt"
)

func main() {
	points := [3]int{1, 2, 3}
	points2 := [...]int{1, 2, 3}
	fmt.Printf("%v, %T\n", points, points)
	fmt.Printf("%v, %T\n", points2, points2)

	var points3 [3]int
	points3[0] = 1
	points3[1] = 1
	points3[2] = 1
	fmt.Printf("%v, %T\n", points3[2], points3) // chuỗi định dạng, print format
	fmt.Println(len(points))


}
