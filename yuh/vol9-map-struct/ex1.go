package main

import "fmt"

type Student struct {
	number int
	name   string
	isMale bool
	subject []string
}

func main() {
	studentYuh := Student{
		number: 1,
		name:   "John",
		isMale: false,
		subject: []string{"Math", "English", "IT"},
	}
	testName := Student{2, "Bruce", false, []string{"Math", "English", "IT"}}
	testName.number = 10
	// define fast
	studentHung := struct {name string}{name: "testName"}
	copyStudentHung := studentHung // sử dụng & biến thành pointer
	copyStudentHung.name = "copy"
	fmt.Println(studentYuh)
	fmt.Println(testName)
	fmt.Println(studentHung)
	fmt.Println(copyStudentHung)
	fmt.Println(studentYuh.name)
}