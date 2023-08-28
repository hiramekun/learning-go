package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	newEmployees := []Employee{
		{"石田三成", "13112"},
		{"徳川家康", "13115"},
	}
	return newEmployees
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "豊臣秀吉",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)
	fmt.Println(m.Description())

	m.Reports = m.FindNewEmployees()
	fmt.Println(m.Employee)
	fmt.Println(m.Reports)
	/**
	12345
	豊臣秀吉 (12345)
	{豊臣秀吉 12345}
	[{石田三成 13112} {徳川家康 13115}]
	*/
}
