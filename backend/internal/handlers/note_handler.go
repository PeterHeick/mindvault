// backend/internal/handlers/note_handler.go
package handlers

import (
	"mindvault/backend/internal/models"
	"mindvault/backend/internal/repositories" // Importer vores nye repository
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// NoteHandler indeholder den logik, der er nødvendig for at håndtere note-requests.
type NoteHandler struct {
	Repo *repositories.NoteRepository
}

// NewNoteHandler opretter en ny instans af NoteHandler.
func NewNoteHandler(repo *repositories.NoteRepository) *NoteHandler {
	return &NoteHandler{Repo: repo}
}

// CreateNote håndterer oprettelsen af en ny note.
func (h *NoteHandler) CreateNote(c *gin.Context) {
	var newNote models.Note

	// Bind den indkommende JSON til vores newNote struct
	if err := c.ShouldBindJSON(&newNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Kald repository'et for at gemme noten i databasen
	if err := h.Repo.CreateNote(&newNote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	// Send den oprettede note (nu med ID og timestamps) tilbage som svar
	c.JSON(http.StatusCreated, newNote)
}

// GetNotes håndterer hentning af alle noter.
func (h *NoteHandler) GetNotes(c *gin.Context) {
	notes, err := h.Repo.GetNotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve notes"})
		return
	}
	c.JSON(http.StatusOK, notes)
}

// GetNoteByID håndterer hentning af en enkelt note.
func (h *NoteHandler) GetNoteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	note, err := h.Repo.GetNoteByID(id)
	if err != nil {
		// pgx.ErrNoRows er den forventede fejl, hvis noten ikke findes.
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}
	c.JSON(http.StatusOK, note)
}


// UpdateNote håndterer opdatering af en eksisterende note.
func (h *NoteHandler) UpdateNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedNote models.Note
	if err := c.ShouldBindJSON(&updatedNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	updatedNote.ID = id // Sæt ID'et fra URL'en

	if err := h.Repo.UpdateNote(&updatedNote); err != nil {
		// Her kunne man specifikt tjekke for pgx.ErrNoRows for at give en 404
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found or failed to update"})
		return
	}

	c.JSON(http.StatusOK, updatedNote)
}

// DeleteNote håndterer sletning af en note.
func (h *NoteHandler) DeleteNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.Repo.DeleteNote(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// 204 No Content er standard HTTP-svar for en succesfuld sletning
	c.Status(http.StatusNoContent)
}