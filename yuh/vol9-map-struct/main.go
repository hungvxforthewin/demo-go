package main

import "fmt"

func main() {
	// kay-value dictionary
	studentNameAgeMap := make(map[string]int)
	studentNameAgeMap = map[string]int {
		"A": 11,
		"B": 22,
		"C": 21,
	}
	studentNameAgeMap["E"] = 21
	delete(studentNameAgeMap, "E")
	_, isExits := studentNameAgeMap["E"]
	// copy pointer
	copySty := studentNameAgeMap
	copySty["A"] = 777
	fmt.Println(studentNameAgeMap)
	fmt.Println(copySty)
	fmt.Println(isExits)
	fmt.Println(len(studentNameAgeMap))
}