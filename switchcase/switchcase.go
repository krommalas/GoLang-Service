package main

import "fmt"

func main() {
	input := 2
	switch input {
	case 1:
		println("Input is 1")
	case 2:
		println("Input is 2")
	case 3:
		println("Input is 3")
	default:
		println("Input is not 1, 2, or 3")
	}

	//switch case name color to rgb
	var color string
	fmt.Println("Enter a color (red, green, blue):")
	fmt.Scanln(&color)
	switch color {
	case "red":
		println("RGB(255, 0, 0)")
	case "green":
		println("RGB(0, 255, 0)")
	case "blue":
		println("RGB(0, 0, 255)")
	default:
		println("Unknown color")
	}
}
