// backend/internal/repositories/note_repository.go
package repositories

import (
	"context"
	"fmt"
	"mindvault/backend/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NoteRepository håndterer al databasekommunikation for noter.
type NoteRepository struct {
	DB *pgxpool.Pool
}

// NewNoteRepository opretter en ny instans af NoteRepository.
func NewNoteRepository(db *pgxpool.Pool) *NoteRepository {
	return &NoteRepository{DB: db}
}

// CreateNote indsætter en ny note i databasen.
func (r *NoteRepository) CreateNote(note *models.Note) error {
	// Definer SQL-query'en. RETURNING henter de værdier, databasen selv genererer.
	sql := `INSERT INTO notes (title, content) VALUES ($1, $2) RETURNING id, created_at, updated_at`

	// Kør query'en. QueryRow bruges, da vi forventer præcis én række tilbage.
	err := r.DB.QueryRow(context.Background(), sql, note.Title, note.Content).Scan(&note.ID, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// GetNotes henter alle noter fra databasen.
func (r *NoteRepository) GetNotes() ([]models.Note, error) {
	sql := `SELECT id, title, content, created_at, updated_at FROM notes ORDER BY created_at DESC`
	rows, err := r.DB.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}

// GetNoteByID henter en enkelt note baseret på ID.
func (r *NoteRepository) GetNoteByID(id int) (*models.Note, error) {
	sql := `SELECT id, title, content, created_at, updated_at FROM notes WHERE id = $1`
	var note models.Note
	err := r.DB.QueryRow(context.Background(), sql, id).Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return nil, err // Kan være 'no rows in result set', hvilket er forventeligt
	}
	return &note, nil
}

// UpdateNote opdaterer en eksisterende note i databasen.
func (r *NoteRepository) UpdateNote(note *models.Note) error {
	sql := `UPDATE notes SET title = $1, content = $2, updated_at = NOW() WHERE id = $3 RETURNING title, content, created_at, updated_at`

	err := r.DB.QueryRow(context.Background(), sql, note.Title, note.Content, note.ID).Scan(&note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return err // Kan fejle, hvis noten ikke findes
	}

	return nil
}

// DeleteNote sletter en note fra databasen.
func (r *NoteRepository) DeleteNote(id int) error {
	sql := `DELETE FROM notes WHERE id = $1`
	
	result, err := r.DB.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("note with id %d not found", id) // Returner en fejl, hvis intet blev slettet
	}

	return nil
}