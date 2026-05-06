package main

import "fmt"

const count = 5

func main() {
	// i := 0
	// for i < 5 {
	// 	println("Value of i:", i)
	// 	i++
	// }

	// for j := 0; j < count; j++ {
	// 	println("Value of j:", j)
	// }
	var input string
	for {
		fmt.Scanf("%s", &input)
		fmt.Println("You entered:", input)
		if input == "exit" {
			fmt.Println("Exiting the loop.")
			break
		}
	}
}
