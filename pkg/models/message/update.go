package message

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

// UpdateRequest.
type UpdateRequest struct {
	// ID.
	ID uuid.UUID

	// ChannelID.
	ChannelID uuid.UUID

	// SenderID.
	SenderID uuid.UUID

	// Content.
	Content *string

	// CreatedAt.
	CreatedAt *time.Time

	// UpdatedAt.
	UpdatedAt *time.Time
}

// UpdateResponse.
type UpdateResponse = Message
