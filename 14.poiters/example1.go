package main

import "fmt"

func main() {
	// Pointer to pointer
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* take the address of var */
	ptr = &a

	/* take the address of ptr using address of operator & */
	pptr = &ptr

	/* take the value using pptr */
	fmt.Printf("Value of a = %d\n", a)
	fmt.Printf("Value available at *ptr = %d\n", *ptr)
	fmt.Printf("Value available at **pptr = %d\n", **pptr)


	// KHI NAO DÙNG &, VÀ *
	// * dùng để khai báo kiểu dữ liệu con trỏ.
	// & dùng để lấy địa chỉ của một biến
	// * liên quan gì với &? trả lời giá trị của con trỏ là địa chỉ của biến mà con trỏ đó chỉ tới	

	// Passing pointer to functions

	// example1 swap number

}