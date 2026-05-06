package main

import "fmt"

func main() {
	var country []string
	country = []string{"Thailand", "Japan", "USA", "Germany", "France"}

	// country = append(country, "Thailand")
	// country = append(country, "Japan")
	// country = append(country, "USA")
	// country = append(country, "Germany")
	// country = append(country, "France")
	fmt.Println(country)

	country = append(country, "Canada", "Australia")
	fmt.Println(country)

	countrySelected := country[:4]
	fmt.Println(countrySelected)

	fmt.Println(country[0])
}
