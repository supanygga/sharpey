package postgres

import (
	"fmt"

	"github.com/supanygga/sharpey/internal/storage/postgres/migrator"
)

// Options.
type Options struct {
	Host     string            `yaml:"host" validate:"required"`
	Port     int               `yaml:"port" default:"5432"`
	User     string            `yaml:"user" validate:"required"`
	Password string            `yaml:"password" validate:"required"`
	Database string            `yaml:"database" validate:"required"`
	Migrate  bool              `yaml:"migrate" default:"true"`
	Migrator *migrator.Options `yaml:"migrator" validate:"required_if=Migrate true"`
}

// connectionString.
func (o *Options) connectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?",
		o.User,
		o.Password,
		o.Host,
		o.Port,
		o.Database,
	)
}
