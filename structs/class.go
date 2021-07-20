package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

// Constructor
func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	employee := Employee{}
	fmt.Printf("%v", employee)

	employee.id = 1
	employee.name = "Alvaro"

	fmt.Printf("%v", employee)
	employee.SetId(100)
	employee.SetName("Javier")
	fmt.Printf("%v\n", employee)

	fmt.Println(employee.GetId())
	fmt.Println(employee.GetName())

	fmt.Println("================")
	employee2 := Employee{
		id:       3,
		name:     "Ragnar",
		vacation: true,
	}
	fmt.Println(employee2)

	fmt.Println("================")
	employee3 := new(Employee)
	fmt.Println(*employee3)

	fmt.Println("================")
	employee4 := NewEmployee(1, "Lagerta", true)
	fmt.Println(*employee4)
}
