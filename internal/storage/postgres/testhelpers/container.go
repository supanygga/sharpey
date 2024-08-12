package testhelpers

import (
	"context"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// New.
func New(ctx context.Context, opts ...WithFunc) (*Container, error) {
	options := defaultOptions()
	options.apply(opts...)

	container, err := postgres.Run(
		ctx,
		options.image,
		postgres.WithUsername(options.user),
		postgres.WithPassword(options.password),
		postgres.WithDatabase(options.database),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}

	port, err := container.MappedPort(ctx, "5432/tcp")
	if err != nil {
		return nil, err
	}

	host, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	connectionString, err := container.ConnectionString(ctx)
	if err != nil {
		return nil, err
	}

	return &Container{
		container:        container,
		host:             host,
		port:             port.Int(),
		user:             options.user,
		password:         options.password,
		database:         options.database,
		connectionString: connectionString,
	}, nil
}

// Host.
func (c *Container) Host() string {
	return c.host
}

// Port.
func (c *Container) Port() int {
	return c.port
}

// User.
func (c *Container) User() string {
	return c.user
}

// Password.
func (c *Container) Password() string {
	return c.password
}

// Database.
func (c *Container) Database() string {
	return c.database
}

// ConnectionString.
func (c *Container) ConnectionString() string {
	return c.connectionString
}

// Terminate.
func (c *Container) Terminate(ctx context.Context) error {
	return c.container.Terminate(ctx)
}
