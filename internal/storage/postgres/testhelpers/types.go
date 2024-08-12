package testhelpers

import "github.com/testcontainers/testcontainers-go/modules/postgres"

// Container
type Container struct {
	container        *postgres.PostgresContainer
	host             string
	port             int
	user             string
	password         string
	database         string
	connectionString string
}
