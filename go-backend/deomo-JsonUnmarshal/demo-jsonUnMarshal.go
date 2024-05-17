package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	e := employee{}
	err := json.Unmarshal([]byte(`{"ID":101,"EmployeeName":"Donny Galaxy","Tel":"0444444444","Email":"donny@mail.com"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e.Tel)
}
