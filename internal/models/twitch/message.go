package twitch

import "time"

// Message.
type Message struct {
	// Channel.
	Channel string

	// User.
	User string

	// Content.
	Content string

	// Timestamp.
	Timestamp time.Time
}
