package main

import (
	"fmt"
	"time"
)

// User ข้อมูลพื้นฐานสำหรับตัวอย่าง
type User struct {
	ID   int
	Name string
}

// ส่งข้อมูล user เข้า channel
func sendUsers(ch chan<- User) {
	users := []User{
		{ID: 1, Name: "Nok"},
		{ID: 2, Name: "Som"},
		{ID: 3, Name: "May"},
	}

	for _, u := range users {
		fmt.Printf("ส่ง user เข้า channel: %+v\n", u)
		ch <- u // ส่ง user ไปยัง channel
		time.Sleep(200 * time.Millisecond)
	}
	close(ch) // ปิด channel เพื่อบอกว่าไม่มีข้อมูลเพิ่ม
}

// รับข้อมูล user จาก channel
func receiveUsers(ch <-chan User, done chan<- bool) {
	for u := range ch { // อ่านจน channel ปิด
		fmt.Printf("รับ user จาก channel: ID=%d Name=%s\n", u.ID, u.Name)
		time.Sleep(300 * time.Millisecond)
	}
	done <- true // แจ้งว่าเสร็จแล้ว
}

func main() {
	userCh := make(chan User)
	done := make(chan bool)

	// สร้าง goroutine สำหรับส่ง user
	go sendUsers(userCh)

	// สร้าง goroutine สำหรับรับ user
	go receiveUsers(userCh, done)

	fmt.Println("รอให้ goroutines ทำงานเสร็จ...")
	<-done // รอจนได้รับสัญญาณจาก receiveUsers
	fmt.Println("เสร็จสิ้นกระบวนการ channel และ goroutine")
}
