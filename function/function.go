package main

import "fmt"

func hello() {
	println("Hello, World!")
}

func plus(value1 int, value2 int) int {
	result := value1 + value2
	// fmt.Println("Result of plus function:", result)
	return result
}

func main() {
	hello()
	result := plus(10, 20)
	fmt.Println("Final result:", result)
}
