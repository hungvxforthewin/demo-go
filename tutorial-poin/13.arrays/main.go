package main

import "fmt"

func main(){
	// Declare array
	var myArray [5] int
	for i := 0; i < 5; i++ {
		myArray[i] = i + 100 /* set element at location i to i + 100 */
	}
	for i := 0; i < len(myArray); i++ {
		fmt.Printf("%d ", myArray[i])
	}
	fmt.Printf("\n")
	fmt.Println("\n")
	// Initializing Arrays
	// var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	// Accessing Array Elements
	var salary int = arr[4]
	fmt.Println(salary)
}

// 	 Không giống đa số các ngôn ngữ khác, Array trong Go không phải là dạng tham chiếu (reference types) mà là dạng tham trị (value types). 
//   Khi gán giá trị nó cho một biến mới thì nó sẽ tạo ra một bản copy của Array cũ, và mọi thay đổi ở Array mới không ảnh hưởng gì đến Array cũ:
