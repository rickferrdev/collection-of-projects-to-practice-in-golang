package users

import (
	"database/sql"
	"log"

	"github.com/rickferrdev/api-basic-v2/internal/database"
)

func CreateUserSchema() *sql.DB {
	conn := database.Connect()

	_, err := conn.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);`)
	if err != nil {
		log.Fatalln(err)
	}

	return conn
}

func ServiceCreateUser(user CreateUserDTO) CreateUserDTO {
	conn := CreateUserSchema()
	defer database.Close(conn)

	_, err := conn.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("User created successfully")

	return user
}

func ServiceObtainUserByID(userID int) User {
	var response User
	conn := CreateUserSchema()
	defer database.Close(conn)

	err := conn.QueryRow("SELECT id, name, email FROM users WHERE id = ?", userID).Scan(&response.ID, &response.Name, &response.Email)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("User retrieved successfully")

	return response
}

func ServiceUpdateUser(user UpdateUserDTO) UpdateUserDTO {
	conn := CreateUserSchema()
	defer database.Close(conn)
	_, err := conn.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("User updated successfully")

	return user
}

func ServiceDeleteUser(userID int) bool {
	conn := CreateUserSchema()
	defer database.Close(conn)

	_, err := conn.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("User deleted successfully")

	return true
}
