package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%s : %d", p.Name, p.Age)
}

//Pointer receiver có thể thay đổi giá trị của đối tượng
func (p *Person) IncreaseAge() {
	p.Age += 1
}

//Value receiver không thay đổi giá trị của của đối tượng
func (p Person) NotReallyIncreaseAge() {
	p.Age += 1 //warning: ineffective assignment to field Person.Age
}
func main() {
	var person1 Person
	person1.Age = 27
	person1.Name = "John"
	// before run function
	fmt.Println(person1.Name, person1.Age)
	// define pointer
	person1.IncreaseAge() //Pointer receiver chấp nhận value
	person1.NotReallyIncreaseAge()
	fmt.Println(person1.Name, person1.Age)

	// define 1 pointer
	var ptr *Person
	ptr = &Person{Name: "human", Age: 18}
	fmt.Println(ptr.Name, ptr.Age)
	ptr.IncreaseAge()
	ptr.NotReallyIncreaseAge() //Value receiver chấp nhận cả pointer
	fmt.Println(ptr.Name, ptr.Age)

}
