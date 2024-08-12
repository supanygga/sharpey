package migrator

import (
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/tern/v2/migrate"
)

// Migrator.
type Migrator struct {
	options  *Options
	logger   *slog.Logger
	conn     *pgx.Conn
	migrator *migrate.Migrator
}
