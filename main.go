package main

import "fmt"

var numberInt, numberInt2 int = 3000, 10000
var msg string = "Hello"

func main() {
	numberfloat := 25.4
	fmt.Println(numberInt)
	fmt.Println(numberInt2)
	fmt.Println(numberfloat)
	fmt.Println(msg)

	fmt.Println(numberInt + numberInt2)
	fmt.Println(float64(numberInt) + numberfloat)
	fmt.Println(msg + "World")
	fmt.Println("my money =", numberInt)
	fmt.Println("your money =", numberInt2)
	fmt.Println("1 inch is equal to", numberfloat, "centimeters")
}
