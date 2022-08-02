package main

import "fmt"

func main() {
	fmt.Println("switch case in golang")
	// 2 type switch
	// type 1
	number := 10
	switch number {
			case 0:
				fmt.Println("incorrect number")
				fmt.Println("incorrect number")
				fmt.Println("incorrect number")
			case 1:
				fmt.Println("incorrect number")
			case 10:
				fmt.Println("correct number")
				break
				fmt.Println("correct number")
				fmt.Println("correct number")
				fmt.Println("correct number")

			default:
				fmt.Println("incorrect number")
	}

	// type 2
	switch {
		case number >= 10:
			fmt.Println("correct number")
		default:
			fmt.Println("incorrect number")
	}
}
