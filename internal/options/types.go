package options

import (
	"github.com/supanygga/sharpey/internal/clients/twitch"
	proc "github.com/supanygga/sharpey/internal/processors/twitch"
	"github.com/supanygga/sharpey/internal/storage/postgres"
)

// Options.
type Options struct {
	// Twitch.
	Twitch *twitch.Options `yaml:"twitch" validate:"required"`

	// Postgres.
	Postgres *postgres.Options `yaml:"postgres" validate:"required"`

	// Processor.
	Processor *proc.Options `yaml:"processor" validate:"required"`
}
