package main

import (
	"log"
	"url-analyser/backend/config"
	"url-analyser/backend/routes"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	r := routes.SetupRouter(db)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
