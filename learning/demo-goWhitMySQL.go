package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func createTable(db *sql.DB) {
	query := `CREATE TABLE user_data (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		status ENUM('active', 'inactive') NOT NULL DEFAULT 'active',
		created_at DATETIME,
		updated_at DATETIME
	);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func insertData(db *sql.DB) {
	var username string
	var password string
	var firstName string
	var lastName string
	var email string
	createAt := time.Now()
	updateAt := time.Now()

	fmt.Print("Enter username: ")
	fmt.Scan(&username)
	fmt.Print("Enter password: ")
	fmt.Scan(&password)
	fmt.Print("Enter first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter email: ")
	fmt.Scan(&email)

	result, err := db.Exec(`INSERT INTO user_data (username, password, first_name, last_name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`, username, password, firstName, lastName, email, createAt, updateAt)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted row with ID: %d\n", id)

}

func deleteData(db *sql.DB) {
	var id int
	fmt.Print("Enter the ID of the user to delete: ")
	fmt.Scan(&id)

	_, err := db.Exec("DELETE FROM user_data WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User deleted successfully!")
}

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
	// query(db)
	// createTable(db)
	// insertData(db)
	deleteData(db)
	defer db.Close()
}
