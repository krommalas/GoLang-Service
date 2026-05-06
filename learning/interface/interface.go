package main

import "fmt"

// Interface สำหรับผู้เล่นในเกม
type Player interface {
	Attack() string
	Defend() string
}

// Struct สำหรับนักรบ
type Warrior struct {
	Name string
}

func (w Warrior) Attack() string {
	return w.Name + " swings a sword!"
}

func (w Warrior) Defend() string {
	return w.Name + " raises a shield!"
}

// Struct สำหรับนักเวทย์
type Mage struct {
	Name string
}

func (m Mage) Attack() string {
	return m.Name + " casts a fireball!"
}

func (m Mage) Defend() string {
	return m.Name + " creates a magic barrier!"
}

func allPlayer(p Player) {
	fmt.Println(p.Attack())
	fmt.Println(p.Defend())
}

func main() {
	var p Player

	// ใช้ Warrior
	p = Warrior{Name: "Aragorn"}
	fmt.Println(p.Attack())
	fmt.Println(p.Defend())

	// ใช้ Mage
	p = Mage{Name: "Gandalf"}
	fmt.Println(p.Attack())
	fmt.Println(p.Defend())

	// ใช้ฟังก์ชันที่รับ Player เป็นพารามิเตอร์
	fmt.Println("\nUsing allPlayer function:")
	allPlayer(Warrior{Name: "Legolas"})
	allPlayer(Mage{Name: "Saruman"})
}
