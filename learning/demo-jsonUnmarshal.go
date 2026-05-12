package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type games struct {
	ID    int
	Name  string
	Price float64
}

func main() {
	g := games{}
	err := json.Unmarshal([]byte(`{"ID":101,"Name":"Dragonball","Price":1400}`), &g)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g)
	fmt.Println(g.Name)
}
