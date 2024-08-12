package message

import (
	"github.com/gofrs/uuid/v5"
)

// CreateRequest.
type CreateRequest struct {
	// ID.
	ID uuid.UUID

	// ChannelID.
	ChannelID uuid.UUID

	// SenderID.
	SenderID uuid.UUID

	// Content.
	Content string
}

// CreateResponse.
type CreateResponse = Message
