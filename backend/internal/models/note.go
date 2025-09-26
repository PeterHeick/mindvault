// backend/internal/models/note.go
package models

import "time"

// Note repræsenterer en enkelt note i systemet.
type Note struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}