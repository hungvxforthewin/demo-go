package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 50, 100}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:5]
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("a: %v, %v, %v\n", a, len(a), cap(a))
	fmt.Printf("b: %v, %v, %v\n", a, len(b), cap(b))
	fmt.Printf("c: %v, %v, %v\n", a, len(c), cap(c))
	fmt.Printf("d: %v, %v, %v\n", a, len(d), cap(d))
	fmt.Printf("e: %v, %v, %v\n", a, len(e), cap(e))

	// len mean current
	// cap tu first item of slice to last item of behind array
	// index of slice đếm từ 1 (khi đến mảng thì phải tính lại 0 (n-1))
	// bản chất slice là pointer của array gốc --> thay đổi 1 thì tất cả cùng thay dổi
}

