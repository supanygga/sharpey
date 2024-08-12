package message

import "github.com/gofrs/uuid/v5"

// GetRequest.
type GetRequest struct {
	// ID.
	ID uuid.UUID
}

// GetResponse.
type GetResponse = Message
