package main

import "os"

func main() {
	data1 := []byte("hello \n world")
	err := os.WriteFile("learning/file/output.txt", data1, 0644)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("learning/file/employees")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data2 := []byte("Nok\nSom\nMay")
	err = os.WriteFile("learning/file/employees.txt", data2, 0644)
	if err != nil {
		panic(err)
	}
}
