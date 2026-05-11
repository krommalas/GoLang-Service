package main

import "fmt"

func process1(c chan string, data string) {
	c <- data
}
func main() {
	ch := make(chan string)
	go process1(ch, "Hello, channel!")
	go process1(ch, "Another message!")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
