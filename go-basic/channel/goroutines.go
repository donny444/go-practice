package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

// func main() {
// 	go f("Hello Go")
// 	go f("message2")
// 	time.Sleep(5 * time.Second)
// }
