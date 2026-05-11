package main

import "fmt"

func add(value1, value2 float64) {
	result := value1 + value2
	fmt.Println("result is = ", result)
}

func loop() {
	for i := 0; i < 5; i++ {
		fmt.Println("loop:", i)
	}
}

func deferLoop() {
	for j := 0; j < 5; j++ {
		defer fmt.Println("defer loop:", j)
	}

}

func main() {
	// defer fmt.Println("Start of main function")
	// defer fmt.Println("End of main function")
	// add(20, 10)
	loop()
	deferLoop()

}
