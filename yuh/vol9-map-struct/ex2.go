package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Humane
	subject []string
}

type Humane struct {
	age  int    `description:"age"`
	name string `description:"name"`
}

func main() {
	hungvx := People{}
	hungvx.age = 27
	hungvx.name = "Hungvx"
	hungvx.subject = []string{"A", "B", "C"}

	fmt.Println(hungvx)
	// get tag
	t := reflect.TypeOf(hungvx)
	fields, _ := t.FieldByName("age")
	fmt.Println(fields)
}
