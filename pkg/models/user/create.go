package user

import (
	"github.com/gofrs/uuid/v5"
)

// CreateRequest.
type CreateRequest struct {
	// ID.
	ID uuid.UUID

	// Name.
	Name string

	// IsBot.
	IsBot bool
}

// CreateResponse.
type CreateResponse = User
