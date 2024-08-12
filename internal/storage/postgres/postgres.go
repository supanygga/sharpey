package postgres

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/supanygga/sharpey/internal/storage/postgres/message"
	"github.com/supanygga/sharpey/internal/storage/postgres/migrator"
	"github.com/supanygga/sharpey/internal/storage/postgres/user"
)

// New.
func New(ctx context.Context, options *Options, logger *slog.Logger) (*Storage, error) {
	if options == nil {
		return nil, errors.New("nil options")
	}

	if logger == nil {
		return nil, errors.New("nil logger")
	}

	conn, err := pgx.Connect(ctx, options.connectionString())
	if err != nil {
		return nil, err
	}

	if options.Migrate {
		migrator, err := migrator.New(ctx, conn, options.Migrator, logger)
		if err != nil {
			return nil, err
		}

		if err := migrator.Migrate(ctx); err != nil {
			return nil, err
		}
	}

	user, err := user.New(conn, logger)
	if err != nil {
		return nil, err
	}

	message, err := message.New(conn, logger)
	if err != nil {
		return nil, err
	}

	return &Storage{
		options: options,
		logger:  logger,
		conn:    conn,
		user:    user,
		message: message,
		closer: func() error {
			return conn.Close(ctx)
		},
	}, nil
}

// Conn.
func (s *Storage) Conn() *pgx.Conn {
	return s.conn
}

// User.
func (s *Storage) User() *user.Storage {
	return s.user
}

// Message.
func (s *Storage) Message() *message.Storage {
	return s.message
}

// Close.
func (s *Storage) Close() error {
	return s.closer()
}
