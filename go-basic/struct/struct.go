package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	employeeList := []employee{}
	employee1 := employee{
		employeeID:   "101",
		employeeName: "Pradoo",
		phone:        "0900000000",
	}
	employee2 := employee{
		employeeID:   "102",
		employeeName: "Prayad",
		phone:        "0900000001",
	}
	employee3 := employee{
		employeeID:   "103",
		employeeName: "Pranee",
		phone:        "0900000002",
	}

	employeeList = append(employeeList, employee1)
	employeeList = append(employeeList, employee2)
	employeeList = append(employeeList, employee3)

	fmt.Println("employee = ", employeeList)
}
