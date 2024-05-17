package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	data, _ := json.Marshal(&employee{101, "Donny Galaxy", "0444444444", "donny@mail.com"})
	fmt.Println(string(data))
}
