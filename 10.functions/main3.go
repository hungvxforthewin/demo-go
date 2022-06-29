package main

import (
	"fmt"
	"math"
)

func main(){
	// Func as value

	// Declare func as variable
	/* declare a function variable */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* use the function */
	fmt.Println(getSquareRoot(9))


	// Func closure
	/* nextNumber is now a function with i as 0 */
	nextNumber := getSequence()
	/* invoke nextNumber to increase i by 1 and return the same */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* create a new sequence and see the result, i is 0 again*/
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}

// Define Anonymous
func getSequence() func() int{
	i:= 0
	//fmt.Print(i)
	return func() int{
		i+=1
		return i
	}
}