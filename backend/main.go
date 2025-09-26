// backend/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"mindvault/backend/internal/handlers"
	"mindvault/backend/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// --- (Konfiguration som før) ---
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL_MODE"))

	// --- (Databaseforbindelse som før) ---
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Kunne ikke oprette forbindelse til databasen: %v\n", err)
	}
	defer pool.Close()
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Kunne ikke pinge databasen: %v\n", err)
	}
	log.Println("Forbindelse til databasen er oprettet succesfuldt!")

	// --- Opret instanser af Repository og Handler ---
	noteRepo := repositories.NewNoteRepository(pool)
	noteHandler := handlers.NewNoteHandler(noteRepo)

	// --- Opsætning af Web Server (Gin) ---
	router := gin.Default()

	// --- API Ruter ---
	v1 := router.Group("/api/v1")
	{
		v1.POST("/notes", noteHandler.CreateNote)
		v1.GET("/notes", noteHandler.GetNotes)
		v1.GET("/notes/:id", noteHandler.GetNoteByID)
		v1.PUT("/notes/:id", noteHandler.UpdateNote)
		v1.DELETE("/notes/:id", noteHandler.DeleteNote)
		// Fremtidige ruter kommer her...
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// --- Start Serveren ---
	log.Printf("MindVault backend starter på http://localhost:%s", appPort)
	if err := router.Run(":" + appPort); err != nil {
		log.Fatalf("Kunne ikke starte serveren: %s", err)
	}
}