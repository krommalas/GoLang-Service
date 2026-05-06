package main

import "fmt"

var number, number2 int = 10, 20
var msg string = "Hello, World!"

func main() {

	numberfloat := 3.14
	fmt.Println("krommalas Mungwicha")
	fmt.Println(number)
	fmt.Println(number2)
	fmt.Println(numberfloat)
	fmt.Println(msg)

	fmt.Println(number + number2)
	fmt.Println(float64(number) + numberfloat)
	fmt.Println(msg + " Welcome to Go programming!")
	fmt.Println("mynumber"+msg, number)
}
