// SQLite-basic/pets/pets.go
package pets

import (
	"database/sql"
	"log"
)

type Pets struct {
	ID   int
	Name string
	Age  int
	Type string
}

// CreateTable creates the pets table if it does not exist
func CreateTable(conn *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS pets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER,
		type TEXT
	);`

	_, err := conn.Exec(query)
	return err
}

// AddPet adds a new pet to the database
func AddPet(conn *sql.DB, pet Pets) error {
	query := `INSERT INTO pets (name, age, type) VALUES (?, ?, ?);`
	_, err := conn.Exec(query, pet.Name, pet.Age, pet.Type)
	return err
}

// GetAllPets retrieves all pets from the database
func GetAllPets(conn *sql.DB) ([]Pets, error) {
	query := `SELECT id, name, age, type FROM pets;`
	rows, err := conn.Query(query)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	defer rows.Close()
	var pets []Pets
	for rows.Next() {
		var pet Pets
		rows.Scan(&pet.ID, &pet.Name, &pet.Age, &pet.Type)
		pets = append(pets, pet)
	}

	return pets, nil
}

// DeletePet removes a pet from the database by ID
func DeletePet(conn *sql.DB, id int) error {
	query := `DELETE FROM pets WHERE id = ?;`
	_, err := conn.Exec(query, id)
	return err
}

// UpdatePet updates an existing pet's information in the database
func UpdatePet(conn *sql.DB, pet Pets) error {
	query := `UPDATE pets SET name = ?, age = ?, type = ? WHERE id = ?;`
	_, err := conn.Exec(query, pet.Name, pet.Age, pet.Type, pet.ID)
	return err
}
