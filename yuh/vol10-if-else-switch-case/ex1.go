package main

import "fmt"

func main() {
	m := map[string]int {
		"a": 1,
		"b": 2,
	}
	if age, isExits := m["a"]; isExits{
		fmt.Println(age)
	}
	a := 10
	guess := 100
	// các toán tử ||, &&
	// if else if
	if a <= guess {
		fmt.Println("10 <= 1000")
	}

}