package user

import (
	"log/slog"

	"github.com/jackc/pgx/v5"
)

// Storage.
type Storage struct {
	conn   *pgx.Conn
	logger *slog.Logger
}
