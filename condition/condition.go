package main

import "fmt"

var score int

func main() {
	//grade condition
	fmt.Println("Enter your score: ")
	fmt.Scanf("%d", &score)
	fmt.Println("Your score is:", score)

	if score >= 90 {
		println("Grade: A")
	} else if score >= 80 {
		println("Grade: B")
	} else if score >= 70 {
		println("Grade: C")
	} else if score >= 60 {
		println("Grade: D")
	} else {
		println("Grade: F")
	}
}
