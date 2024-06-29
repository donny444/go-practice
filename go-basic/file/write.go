package main

import "os"

func main() {
	data1 := []byte("Test File Writing\nin Go Language")
	err := os.WriteFile("file/data.txt", data1, 0644)
	if err != nil {
		panic(err)
	}
	f, ferr := os.Create("file/customerName.txt")
	if ferr != nil {
		panic(ferr)
	}

	defer f.Close()

	data2 := []byte("John\nWidth")
	os.WriteFile("file/employeeName.txt", data2, 0644)
}
