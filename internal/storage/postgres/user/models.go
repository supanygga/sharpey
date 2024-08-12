package user

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

// create.
type create struct {
	ID    uuid.UUID `db:"id"`
	Name  string    `db:"name"`
	IsBot bool      `db:"is_bot"`
}

// update.
type update struct {
	ID        uuid.UUID  `db:"id"`
	Name      *string    `db:"name"`
	IsBot     *bool      `db:"is_bot"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

func idk() {
	uuid.NewV7()
}
