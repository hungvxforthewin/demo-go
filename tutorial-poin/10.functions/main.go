package main

import "fmt"

func main()  {
	/* local variable definition */
	var a int = 100
	var b int = 200
	var ret int

	/* calling a function to get max value */
	ret = max(a, b)

	fmt.Printf( "Max value is : %d\n", ret )

	/* calling func return multi values */
	x, y := swap("A", "B")
	fmt.Println(x, y)
}

// Declare function
func max(num1 int, num2 int)  int{
	var result int
	if(num1 > num2){
		result = num1
	}else{
		result = num2
	}
	return result
}

// Func return multi values
func swap(x, y string) (string, string){
	return y, x
}
