package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/weaver/api/internal/config"
	"github.com/weaver/api/internal/database"
	"github.com/weaver/api/internal/handlers"
)

func main() {
	// Load environment variables
	_ = godotenv.Load()

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := database.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Seed demo data
	if err := database.SeedDemoData(db); err != nil {
		log.Printf("Warning: Failed to seed demo data: %v", err)
	}

	// Create Gin router
	router := gin.Default()

	// Setup routes (includes middleware)
	handlers.SetupRoutes(router, db)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Starting server on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
