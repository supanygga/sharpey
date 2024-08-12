package user

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

// User.
type User struct {
	// ID.
	ID uuid.UUID

	// Name.
	Name string

	// IsBot.
	IsBot bool

	// CreatedAt.
	CreatedAt time.Time

	// UpdatedAt.
	UpdatedAt time.Time
}
