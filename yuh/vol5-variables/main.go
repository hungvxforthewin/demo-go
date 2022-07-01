package main

import "fmt"

// global variables
var global_variable = "global"
// export global scope
var Number int = 10
// block scope
var (
	a int = 1
	b int = 2
	m string = "abc"
)

func main() {

	var x int
	x = 1
	var y int = 10
	z := 3.14
	var str string = string(x)
	// shadow
	var a int = 30
	// convert type
	var t float32 = float32(x)
	fmt.Println(y, z, a, t)
	fmt.Printf("x: %d\n", x)
	fmt.Printf("x: %d\n", x)
	fmt.Printf("y: %v, %T", y, y)
	fmt.Printf("z: %v, %T", z, z)
	fmt.Printf("global_variable: %s, %T", global_variable, global_variable)
	fmt.Printf("\nstr: %v, %T", str, str)


}
/*
	1. Định nghĩa biến trong ngôn ngữ lập trình
	2. Cú pháp và khai báo biên, khởi tạo giá trị
	3. Global and block scope
	4. Shadow
	5. Declared and not used (Compiler error)
	6. Export global scope
	7. Naming conventions
		camel(lạc đà)
		helloString
	8. Convert type
		float32 --> int: bỏ .
		int --> string (follow unicode): sử dụng string convert
*/