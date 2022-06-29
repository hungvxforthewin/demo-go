package main

import "fmt"

func main() {
	fmt.Println("Operator")
	// Arithmetic Operator
	var a1, b1 = 10, 20
	fmt.Printf("a + b is %d\n", a1+b1)
	fmt.Printf("a - b is %d\n", a1-b1)
	fmt.Printf("a * b is %d\n", a1*b1)
	fmt.Printf("a / b is %d\n", a1/b1)
	fmt.Printf("a % b is %d\n", a1%b1)
	a1++
	fmt.Printf("a ++ is %d\n", a1)
	b1--
	fmt.Printf("b -- is %d\n", b1)

	// Relational Operators
	/*
		==
		!=
		<
		>
		<=
		>=
	*/
	var a int = 21
	var b int = 10

	if a == b {
		fmt.Printf("Line 1 - a is equal to b\n")
	} else {
		fmt.Printf("Line 1 - a is not equal to b\n")
	}
	if a < b {
		fmt.Printf("Line 2 - a is less than b\n")
	} else {
		fmt.Printf("Line 2 - a is not less than b\n")
	}
	if a > b {
		fmt.Printf("Line 3 - a is greater than b\n")
	} else {
		fmt.Printf("Line 3 - a is not greater than b\n")
	}

	/* Lets change value of a and b */
	a = 5
	b = 20
	if a <= b {
		fmt.Printf("Line 4 - a is either less than or equal to  b\n")
	}
	if b >= a {
		fmt.Printf("Line 5 - b is either greater than  or equal to b\n")
	}

	// Logical Operators

	/*
		&&
		||
		!
	*/

	var a2 bool = true
	var b2 bool = false
	if a2 && b2 {
		fmt.Printf("Line 1 - Condition is true\n")
	}
	if a2 || b2 {
		fmt.Printf("Line 2 - Condition is true\n")
	}

	/* lets change the value of  a and b */
	a2 = false
	b2 = true
	if a2 && b2 {
		fmt.Printf("Line 3 - Condition is true\n")
	} else {
		fmt.Printf("Line 3 - Condition is not true\n")
	}
	if !(a2 && b2) {
		fmt.Printf("Line 4 - Condition is true\n")
	}

	// Assignment Operators
	/*
		=
		+=
		-=
		*=
		/=
		%=
	*/

	// Miscellaneous Operators
	/*
		& return the address of a variable
		* pointer to a variable
	*/
	var x int = 4
	var y int32
	var z float32
	var ptr *int

	/* example of type operator */
	fmt.Printf("Line 1 - Type of variable x = %T\n", x)
	fmt.Printf("Line 2 - Type of variable y = %T\n", y)
	fmt.Printf("Line 3 - Type of variable z= %T\n", z)
	fmt.Printf("Line 4 - Type of variable ptr= %T\n", ptr)

	ptr = &x
	x = 10
	fmt.Printf("value of x is: %d\n", x)
	fmt.Printf("value of ptr is: %d\n", ptr)
	fmt.Printf("*ptr is: %d", *ptr)

}
