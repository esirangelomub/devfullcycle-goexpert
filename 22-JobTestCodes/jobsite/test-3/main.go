package main

type Employee struct {
	FirstName string
	LastName  string
}

type Manager struct {
	Employee
	Salary int
}

func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

func main() {
	m := Manager{
		Employee: Employee{
			FirstName: "John",
			LastName:  "Doe",
		},
		Salary: 3000,
	}

	// Right
	println(m.FullName())
	println(m.Employee.FullName())
	println(m.FirstName + " " + m.LastName)

	// Wrong
	println(m.Employee.FullName) // show the address of the function
	//println(m.Field(0).FirstName) // not working
}
