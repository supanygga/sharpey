package postgres

import (
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/supanygga/sharpey/internal/storage/postgres/message"
	"github.com/supanygga/sharpey/internal/storage/postgres/user"
)

// Storage.
type Storage struct {
	options *Options
	logger  *slog.Logger
	conn    *pgx.Conn
	user    *user.Storage
	message *message.Storage
	closer  func() error
}
