package test

import (
	"database/sql"
	"testing"

	"github.com/rickferrdev/sqlite-basic/pets"
	_ "modernc.org/sqlite"
)

func TestAddPet(t *testing.T) {
	conn := MemoryDBTestSetup(t)
	defer conn.Close()

	pet := pets.Pets{Name: "Buddy", Age: 3, Type: "Dog"}

	err := pets.AddPet(conn, pet)
	if err != nil {
		t.Fatalf("Failed to add pet: %v", err)
	}
}

func CloseDB(t *testing.T, db *sql.DB) {
	if err := db.Close(); err != nil {
		t.Fatalf("Failed to close database: %v", err)
	}
}

func TestGetAllPets(t *testing.T) {
	conn := MemoryDBTestSetup(t)
	defer conn.Close()
	pet1 := pets.Pets{Name: "Buddy", Age: 3, Type: "Dog"}
	pet2 := pets.Pets{Name: "Mittens", Age: 2, Type: "Cat"}
	err := pets.AddPet(conn, pet1)
	if err != nil {
		t.Fatalf("Failed to add pet1: %v", err)
	}
	err = pets.AddPet(conn, pet2)
	if err != nil {
		t.Fatalf("Failed to add pet2: %v", err)
	}
	petsList, err := pets.GetAllPets(conn)
	if err != nil {
		t.Fatalf("Failed to get all pets: %v", err)
	}
	if len(petsList) >= 2 {
		t.Logf("Successfully retrieved pets: %+v", petsList)
	}
}

func TestDeletePet(t *testing.T) {
	conn := MemoryDBTestSetup(t)
	defer conn.Close()

	pet := pets.Pets{Name: "Buddy", Age: 3, Type: "Dog"}
	err := pets.AddPet(conn, pet)
	if err != nil {
		t.Fatalf("Failed to add pet: %v", err)
	}

	err = pets.DeletePet(conn, 1)
	if err != nil {
		t.Fatalf("Failed to delete pet: %v", err)
	}
}

func MemoryDBTestSetup(t *testing.T) *sql.DB {
	conn, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open in-memory database: %v", err)
	}

	err = pets.CreateTable(conn)
	if err != nil {
		t.Fatalf("Failed to create pets table: %v", err)
	}

	return conn
}
