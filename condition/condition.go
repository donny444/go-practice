package main

import "fmt"

var score int

func main() {
	fmt.Println("grade calculator")
	fmt.Scanf("%d", &score)
	fmt.Println("score = ", score)
	if score >= 80 {
		fmt.Println("You got A grade")
	} else if score >= 70 {
		fmt.Println("You got B grade")
	} else if score >= 60 {
		fmt.Println("You got C grade")
	} else if score >= 50 {
		fmt.Println("You got D grade")
	} else {
		fmt.Println("You got F grade")
	}
}
