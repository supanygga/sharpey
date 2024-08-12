package user

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

// UpdateRequest.
type UpdateRequest struct {
	// ID.
	ID uuid.UUID

	// Name.
	Name *string

	// IsBot.
	IsBot *bool

	// CreatedAt.
	CreatedAt *time.Time

	// UpdatedAt.
	UpdatedAt *time.Time
}

// UpdateResponse.
type UpdateResponse = User
