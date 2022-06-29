package main

import "fmt"

func main()  {
	/* local variable definition */
	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, value of a : %d\n", a )
	fmt.Printf("Before swap, value of b : %d\n", b )

	/* calling a function to swap the values */
	//swap_1(a, b)
	swap_2(&a, &b)

	fmt.Printf("After swap, value of a : %d\n", a )
	fmt.Printf("After swap, value of b : %d\n", b )
}

func swap_1(x int, y int) int{
	var temp int
	temp = x
	x = y
	y = temp
	return  temp
}

func swap_2(x *int, y *int) int{
	var temp int
	temp = *x
	*x = *y
	*y = temp
	return  temp
}