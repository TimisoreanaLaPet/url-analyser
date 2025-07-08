package main

import (
	"log"
	"url-analyser/backend/config"
	"url-analyser/backend/models"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	err = db.AutoMigrate(&models.URLAnalysis{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Tables created successfully!")
}
