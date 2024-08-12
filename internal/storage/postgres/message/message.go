package message

import (
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

// New.
func New(conn *pgx.Conn, logger *slog.Logger) (*Storage, error) {
	if conn == nil {
		return nil, errors.New("nil pgx connection")
	}

	if logger == nil {
		return nil, errors.New("nil logger")
	}

	return &Storage{
		conn:   conn,
		logger: logger,
	}, nil
}
