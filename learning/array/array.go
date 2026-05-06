package main

import "fmt"

var products [5]string
var price [5]float64

func main() {
	products[0] = "Laptop"
	products[1] = "Mouse"
	products[2] = "Keyboard"
	products[3] = "Monitor"
	products[4] = "Headphones"

	// price[0] = 999.99
	// price[1] = 25.50
	// price[2] = 45.00
	// price[3] = 199.99
	// price[4] = 89.99

	price = [5]float64{999.99, 25.50, 45.00, 199.99, 89.99}

	fmt.Println(products)
	fmt.Println(price)
	fmt.Println(products[3])
	fmt.Println(price[3])
}
