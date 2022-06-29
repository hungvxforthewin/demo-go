package main

import "fmt"

func main() {
	fmt.Println("Condition")
	/*
		if else, if, if else if, nested if statements
	*/
	/*
		switch case
		Expression Switch
		Type Switch
	*/
	/* local variable definition */
	var grade string = "B"
	var marks int = 90
	switch marks {
	case 90:
		grade = "A"
		break
	case 80:
		grade = "B"
		break
	case 50, 60, 70:
		grade = "C"
		break
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("Excellent!\n")
	case grade == "B", grade == "C":
		fmt.Printf("Well done\n")
	case grade == "D":
		fmt.Printf("You passed\n")
	case grade == "F":
		fmt.Printf("Better try again\n")
	default:
		fmt.Printf("Invalid grade\n")
	}
	fmt.Printf("Your grade is  %s\n", grade)

	// Type Switch
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("type of x :%T",i)
	case int:
		fmt.Printf("x is int")
	case float64:
		fmt.Printf("x is float64")
	case func(int) float64:
		fmt.Printf("x is func(int)")
	case bool, string:
		fmt.Printf("x is bool or string")
	default:
		fmt.Printf("don't know the type")
	}
}
