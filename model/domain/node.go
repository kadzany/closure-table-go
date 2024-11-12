package domain

import (
	"database/sql"
	"github.com/google/uuid"
)

type Node struct {
	ID          uuid.UUID      `db:"id" json:"id"`
	Title       string         `db:"title" json:"title"`
	Type        string         `db:"type" json:"type"`
	Description sql.NullString `db:"description,omitempty" json:"description,omitempty"`
	CreatedAt   sql.NullTime   `db:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   sql.NullTime   `db:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt   sql.NullTime   `db:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}
