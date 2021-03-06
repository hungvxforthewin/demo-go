package main

import "fmt"

func main() {
	/* create a slice */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* print the original slice */
	fmt.Println("numbers ==", numbers)

	/* print the sub slice starting from index 1(included) to index 4(excluded)*/
	// trí 1 (low) đến phần tử ở vị trí 3 (high - 1)
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* missing lower bound implies 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* missing upper bound implies len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* print the sub slice starting from index 0(included) to index 2(excluded) */
	number2 := numbers[:2] // cap dc tính từ vị trí thứ 0
	printSlice(number2)

	/* print the sub slice starting from index 2(included) to index 5(excluded) */
	number3 := numbers[2:5] // cap dc tính từ vị trí thứ 2 (thứ 0)
	printSlice(number3)

	println("DEMO APPEND")
	demoAppend()
}

func printSlice(x []int){
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x),x)
}

// append(), copy()

func demoAppend(){
	var numbers []int
	printSlice(numbers)

	/* append allows nil slice */
	numbers = append(numbers, 0)
	printSlice(numbers)
	/* add one element to slice*/
	numbers = append(numbers, 1)
	printSlice(numbers)
	/* add more than one element at a time*/
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)
}