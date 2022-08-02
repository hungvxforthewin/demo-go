package main

import "fmt"

func main() {
	fmt.Println("loop in go")
	for i := 0; i < 10; i++ {
		if i == 0 {
			continue
		}
		if i == 8 {
			break
		}
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
