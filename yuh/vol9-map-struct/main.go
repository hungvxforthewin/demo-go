package main

import "fmt"

func main() {
	studentNameAgeMap := make(map[string]int)
	studentNameAgeMap = map[string]int {
		"A": 11,
		"B": 22,
		"C": 21,
	}
	studentNameAgeMap["E"] = 21
	delete(studentNameAgeMap, "E")
	_, isExits := studentNameAgeMap["E"]
	fmt.Println(studentNameAgeMap)
	fmt.Println(isExits)
}