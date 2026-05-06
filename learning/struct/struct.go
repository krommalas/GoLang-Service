package main

import "fmt"

type employee struct {
	employeeID string
	name       string
	age        int
}

func main() {
	employee1 := employee{
		employeeID: "E001",
		name:       "John Doe",
		age:        30,
	}
	fmt.Println("Employee 1:", employee1)

	employeeList := [3]employee{}
	employeeList[0] = employee1
	employeeList[1] = employee{
		employeeID: "E002",
		name:       "Jane Smith",
		age:        28,
	}
	employeeList[2] = employee{
		employeeID: "E003",
		name:       "Bob Johnson",
		age:        35,
	}
	fmt.Println("Employee List:", employeeList)
}
