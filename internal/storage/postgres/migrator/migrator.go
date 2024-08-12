package migrator

import (
	"context"
	"errors"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/tern/v2/migrate"
)

// New.
func New(ctx context.Context, conn *pgx.Conn, options *Options, logger *slog.Logger) (*Migrator, error) {
	if conn == nil {
		return nil, errors.New("nil pgx connection")
	}

	if options == nil {
		return nil, errors.New("nil options")
	}

	if logger == nil {
		return nil, errors.New("nil logger")
	}

	migrator, err := migrate.NewMigrator(ctx, conn, options.Table)
	if err != nil {
		return nil, err
	}

	return &Migrator{
		conn:     conn,
		options:  options,
		logger:   logger,
		migrator: migrator,
	}, nil
}

// Migrate.
func (m *Migrator) Migrate(ctx context.Context) error {
	currentVersion, err := m.migrator.GetCurrentVersion(ctx)
	if err != nil {
		return err
	}

	if currentVersion == m.options.Version {
		return nil
	}

	if err := m.migrator.LoadMigrations(os.DirFS(m.options.Location)); err != nil {
		return err
	}

	return m.migrator.MigrateTo(ctx, m.options.Version)
}
