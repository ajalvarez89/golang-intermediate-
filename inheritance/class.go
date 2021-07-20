package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	//campo anonimo
	Person
	Employee
}

func (employee FullTimeEmployee) String() string {
	return fmt.Sprintf("\nid: %d, name: %s, age: %d ", employee.id, employee.name, employee.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Thor"
	ftEmployee.age = 32
	ftEmployee.id = 2

	fmt.Println(ftEmployee)
}
