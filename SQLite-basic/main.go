package main

import (
	"log"

	"github.com/rickferrdev/sqlite-basic/connection"
	"github.com/rickferrdev/sqlite-basic/pets"
)

func main() {
	conn := connection.Connect()
	defer connection.Close(conn)
	pets.CreateTable(conn)

	pets.AddPet(conn, pets.Pets{
		ID:   1,
		Name: "Bobby",
		Age:  3,
		Type: "Poodle",
	})

	log.Println("Pet added successfully")

	pets.AddPet(conn, pets.Pets{
		ID:   2,
		Name: "Mimi",
		Age:  1,
		Type: "Siamese Cat",
	})

	log.Println("Pet added successfully")
	allPets, err := pets.GetAllPets(conn)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("All Pets:", allPets)
	err = pets.DeletePet(conn, 1)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Pet deleted successfully")

	allPets, err = pets.GetAllPets(conn)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("All Pets after deletion:", allPets)
}
