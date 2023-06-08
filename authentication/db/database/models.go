// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Bio       sql.NullString
	Password  string
	Isadmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
