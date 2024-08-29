package tests

import (
	"context"
	"os"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresTestContainer struct {
	Context context.Context
	testcontainers.Container
}

func NewPostgresTestContainer(ctx context.Context) *PostgresTestContainer {
	return &PostgresTestContainer{
		Context: ctx,
	}
}

func (c *PostgresTestContainer) Setup() error {

	port := os.Getenv("POSTGRES_PORT")
	db := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")

	req := testcontainers.ContainerRequest{
		Image:        "postgres:12",
		ExposedPorts: []string{port + "/tcp"},
		WaitingFor:   wait.ForLog("database system is ready to accept connections"),
		Env: map[string]string{
			"POSTGRES_USER":     user,
			"POSTGRES_PASSWORD": password,
			"POSTGRES_DB":       db,
		},
	}

	postgresC, err := testcontainers.GenericContainer(c.Context, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return err
	}

	c.Container = postgresC

	return nil
}

func (c *PostgresTestContainer) Terminate() error {
	return c.Container.Terminate(c.Context)
}
