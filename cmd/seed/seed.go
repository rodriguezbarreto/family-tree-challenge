package main

import (
	"family-tree-challenge/cmd/configs"
	"family-tree-challenge/internal/infra/database"
	"log"
)

func main() {
	println("Seeding database...")
	db := database.PostgresConnection()
	err := configs.SeedDatabase(db)
	if err != nil {
		log.Fatalf("Error seeding database: %s", err)
		return
	}
	println("Database successfully seeded ✔")


}
