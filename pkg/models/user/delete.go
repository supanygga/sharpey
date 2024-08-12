package user

import "github.com/gofrs/uuid/v5"

// DeleteRequest.
type DeleteRequest struct {
	// ID.
	ID uuid.UUID
}

// DeleteResponse.
type DeleteResponse = User
