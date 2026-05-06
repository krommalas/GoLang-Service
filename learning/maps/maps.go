package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("Product prices:", product)

	//add products and their prices
	product["Laptop"] = 999.99
	product["Mouse"] = 25.50
	product["Keyboard"] = 45.00
	product["Monitor"] = 199.99
	product["Headphones"] = 89.99

	fmt.Println("Product prices:", product)

	//delete a product
	delete(product, "Mouse")
	fmt.Println("Product prices after deletion:", product)

	//update a product price
	product["Laptop"] = 899.99
	fmt.Println("Product prices after update:", product)

	//access a product price
	fmt.Println("Price of Laptop:", product["Laptop"])

	//calculate total price of Laptop and Monitor
	value1 := product["Laptop"]
	value2 := product["Monitor"]
	fmt.Println("Total price of Laptop and Monitor:", value1+value2)

	food := map[string]string{"Pizza": "Italian", "Sushi": "Japanese", "Tacos": "Mexican"}
	fmt.Println("Food types:", food)
}
