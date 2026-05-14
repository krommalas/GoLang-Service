package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func query(db *sql.DB) {
	var (
		id      int
		name    string
		price   float64
		company string
	)
	query := "SELECT id, name, price, company FROM games_data WHERE id = ?"
	if err := db.QueryRow(query, 1).Scan(&id, &name, &price, &company); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %d, Name: %s, Price: %.2f, Company: %s\n", id, name, price, company)
}

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", "root:Admin@2026@tcp(127.0.0.1:3306)/gamesdb")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to the database successfully!")
	}
	fmt.Println(db)
	query(db)
	defer db.Close()
}
