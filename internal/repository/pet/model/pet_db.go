package model

import (
	"github.com/google/uuid"
	"time"
)

type PetDB struct {
	ID        uuid.UUID `db:"id"`
	Status    string    `db:"status"`
	Name      string    `db:"name"`
	Passport  string    `db:"passport"` // json
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Version   time.Time `db:"version"`
}
