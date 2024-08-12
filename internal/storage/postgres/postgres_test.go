package postgres_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/supanygga/sharpey/internal/storage/postgres"
	"github.com/supanygga/sharpey/internal/storage/postgres/migrator"
	"github.com/supanygga/sharpey/internal/storage/postgres/testhelpers"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	logger := slog.Default()
	ctx := context.Background()
	container, err := testhelpers.New(ctx)
	require.NoError(err)
	assert.NotNil(container)

	t.Cleanup(func() { container.Terminate(ctx) })

	pg, err := postgres.New(ctx, &postgres.Options{
		Host:     container.Host(),
		Port:     container.Port(),
		User:     container.User(),
		Password: container.Password(),
		Database: container.Database(),
		Migrate:  true,
		Migrator: &migrator.Options{
			Table:    "schema_version",
			Location: "../../../migrations",
			Version:  1,
		},
	}, logger)
	require.NoError(err)
	assert.NotNil(pg)

}
