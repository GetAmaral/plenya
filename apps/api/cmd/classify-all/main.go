package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

func main() {
	// Get DB connection string from env
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "5432"
	}
	if dbUser == "" {
		dbUser = "plenya_user"
	}
	if dbPassword == "" {
		dbPassword = "plenya_password"
	}
	if dbName == "" {
		dbName = "plenya_db"
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database successfully")

	// Create service
	batchService := services.NewLabResultBatchService(db)

	// Get all batches
	var batches []models.LabResultBatch
	if err := db.Find(&batches).Error; err != nil {
		log.Fatalf("Failed to get batches: %v", err)
	}

	log.Printf("Found %d batches to classify", len(batches))

	// Classify each batch
	successCount := 0
	errorCount := 0

	for i, batch := range batches {
		log.Printf("[%d/%d] Classifying batch %s...", i+1, len(batches), batch.ID)

		if err := batchService.ClassifyBatchResults(batch.ID); err != nil {
			log.Printf("  ❌ Error: %v", err)
			errorCount++
		} else {
			log.Printf("  ✅ Success")
			successCount++
		}
	}

	log.Println("==========================================")
	log.Printf("Classification complete!")
	log.Printf("  Total batches: %d", len(batches))
	log.Printf("  Successful: %d", successCount)
	log.Printf("  Errors: %d", errorCount)
	log.Println("==========================================")
}
