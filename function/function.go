package main

import "fmt"

var name string

func hello() {
	fmt.Println("Hello Go")
}

func plus(value1 int, value2 int) int {
	return value1 + value2
}

func plus3value(value1, value2, value3 int) int {
	return value1 + value2 + value3
}

func goodbye(name string) {
	fmt.Println("Goodbye", name)
}

func main() {
	hello()
	result := plus(5, 5)
	fmt.Println("result=", result)

	result2 := plus3value(5, 5, 10)
	fmt.Println("result2=", result2)
	fmt.Println("Please enter your name")
	fmt.Scanf("%s", &name)
	goodbye(name)
}
