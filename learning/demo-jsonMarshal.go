package main

import (
	"encoding/json"
	"fmt"
)

type games struct {
	ID    int
	Name  string
	Price float64
}

func main() {
	data, _ := json.Marshal(&games{101, "Dragonball", 1400})
	fmt.Println(string(data))
}
