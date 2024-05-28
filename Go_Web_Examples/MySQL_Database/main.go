// main.go

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Open connection to the database
	dbConnect()

	// Drop table if it exists
	dropTable()
	// Create a new table
	createTable()

	// Add users
	addUser("Big Mike", "SECRET")
	addUser("Little John", "GUESS")
	addUser("Tiny Tim", "NOPE")
	fmt.Println()

	// Query a single user
	queryUser(2)

	// Query all users
	queryAllUsers()
}

// Open connection to the database
func dbConnect() {
	var err error
	db, err = sql.Open("mysql", "mike:5454160s@(127.0.0.1:3306)/my_database?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Database opened successfully.\n")
}

// Drop table if it exists
func dropTable() {
	query := `DROP TABLE IF EXISTS users;`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

// Create a new table
func createTable() {
	query := `
            CREATE TABLE IF NOT EXISTS users (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                PRIMARY KEY (id)
            );`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
	log.Printf("Table created successfully.\n\n")
}

// Insert a user
func addUser(username string, password string) {
	result, err := db.Exec(`INSERT INTO users (username, password) VALUES (?, ?)`, username, password)
	if err != nil {
		log.Fatal(err)
	}

	id, _ := result.LastInsertId()
	log.Printf("Record inserted. Id: %v\n", id)
}

// Query a single user
func queryUser(user_id int) {
	var (
		id       int
		username string
		password string
	)

	query := "SELECT id, username, password FROM users WHERE id = ?"
	if err := db.QueryRow(query, user_id).Scan(&id, &username, &password); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Query a single user:\n")
	fmt.Println("ID:\tUsername:\tPassword:")
	fmt.Printf("%v\t%v\t%v\n\n", id, username, password)
}

// Query all users
func queryAllUsers() {
	type User struct {
		id       int
		username string
		password string
	}

	rows, err := db.Query("SELECT id, username, password FROM users;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User

		err := rows.Scan(&user.id, &user.username, &user.password)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Query all users:\n")
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}
