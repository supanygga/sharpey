package message

import (
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/supanygga/sharpey/pkg/models/user"
)

// Message.
type Message struct {
	// ID.
	ID uuid.UUID

	// Channel.
	Channel *user.User

	// Sender.
	Sender *user.User

	// Content.
	Content string

	// CreatedAt.
	CreatedAt time.Time

	// UpdatedAt.
	UpdatedAt time.Time
}
