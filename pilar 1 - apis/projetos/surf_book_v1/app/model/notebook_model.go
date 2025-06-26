package model

import (
	"time"

	"github.com/google/uuid"
)

type NotebookEntity struct {
	NotebookID  uuid.UUID  `json:"notebook_id"`
	UserID      uuid.UUID  `json:"user_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Icon        string     `json:"icon"`
	Image       string     `json:"image"`
	DeletedAt   *time.Time `json:"deleted_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedAt   time.Time  `json:"created_at"`
}

type CreateNotebookDTO struct {
	UserID uuid.UUID

	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Image       string `json:"image"`
}

type FindNotebookFromUserDTO struct {
	UserID     uuid.UUID
	NotebookID uuid.UUID
}

type ListNotebookFromUserDTO struct {
	UserID uuid.UUID
}

type DeleteNotebookDTO struct {
	NotebookID uuid.UUID
	UserID     uuid.UUID
}

type UpdateNotebookDTO struct {
	UserID      uuid.UUID
	NotebookID  uuid.UUID
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Image       string `json:"image"`
}
