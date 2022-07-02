package main

import (
	"fmt"
)

const c int = 10

const (
	red   = 0
	black = 1
	green = "green"
)

const (
	//errorDefault = iota
	_ = iota + 5 // ko ton ram
	red1
	black1
	green1
)

func main() {
	const a int16 = 10 + 1
	const b = "Go"
	const c = 22
	// const d = math.Sqrt(4)
	//a = 64
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", red, red)
	fmt.Printf("%v, %T\n", black, black)
	fmt.Printf("%v, %T\n", green, green)

	fmt.Printf("%v, %T\n", red1, red1)
	fmt.Printf("%v, %T\n", black1, black1)
	fmt.Printf("%v, %T\n", green1, green1)
	//fmt.Printf("%v, %T\n", d, d)
}

/*
1. Đĩnh nghĩa hằng số
2. Khai báo hăng số
3. Đặc điểm của hằng số
		phỉa có value (cần phải tính toán trong lúc biên dịch)
		ko assign lại data
		assign data --> type
		shadow
4. Kiểu dữ liệu Enum
5. khởi tạo value enum  = iota
	iota phụ thuộc vào scope ()

*/
