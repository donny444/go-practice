package main

import "os"

func main() {
	data1 := []byte("Hi\n Go")
	err := os.WriteFile("file/data.txt", data1, 0644)
	if err != nil {
		panic(err)
	}
	f, ferr := os.Create("file/employeeName")
	if ferr != nil {
		panic(ferr)
	}

	defer f.Close()

	data2 := []byte("Sira\n Manee")
	os.WriteFile("file/employeeName.txt", data2, 0644)
}
