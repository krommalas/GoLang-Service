package main

import "fmt"

func zeroValue(ivalue int) {
	ivalue = 0
	fmt.Println("Zero value of int:", ivalue)
}

func zeroPointer(ipointer *int) {
	*ipointer = 0
	// fmt.Println("Zero value of int through pointer:", *ipointer)
}

func main() {

	i := 10
	fmt.Println("i =", i)

	zeroValue(i)
	fmt.Println("i after zeroValue function call:", i)

	zeroPointer(&i)
	fmt.Println("i after zeroPointer function call:", i)
	fmt.Println("Address of i:", &i)
}
