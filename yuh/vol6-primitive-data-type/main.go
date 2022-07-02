package main

import (
	"fmt"
)

func main() {
	var a bool = true
	var b bool // default false
	var c bool
	c = 1 == 2
	// float
	var a1 float64 = 3.14
	a1 = 13.7e72
	// string, byte
	var str string = "hello world"
	var bstr []byte = []byte(str)
	var cstr byte = 'H'
	// run
	var h rune = 'á'

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", a1, a1)
	fmt.Printf("%v, %T\n", bstr, bstr)
	fmt.Printf("%v, %T\n", cstr, cstr)
	fmt.Printf("%v, %T\n", string(cstr), cstr)
	fmt.Printf("%v, %T\n", string(h), h)
}

/*
	Primitive data type
	1. Boolean
	2. Number
		integer
			signed integer (có dấu) int8, int16, int32, int64
			unsigned integer		uint8, uint16, uint32, uint64
			--> khoảng chứa giá trị
			--> các toán tử (nhị phân &, |, ^, &^)
			&	1 + 1 = 1
				1 + 0 = 0 + 1 = 0
				0 + 0 = 0
			|
				1 + 0 = 1
				0 + 1 = 1
				1 + 1 = 1
				0 + 0 = 0
			^
				2 bit giống nhau = 0
			&^
				2 số 0 = 1
				còn lại = 0
			--> phép dịch bit

		float
			float64 (64bit để lưu trữ)
			float32
		complex
			complex64
			complex128

	3. Text
		String
			cộng string
		Byte --> UTF-8 (chữ cơ bản bảng mã unicode)
		Rune --> UTF-32

*/
