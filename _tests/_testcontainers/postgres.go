package _testcontainers

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"sync"
	"testing"
)

var postgresRunning bool
var pgLock = &sync.Mutex{}
var pgWg = &countWaitGroup{}
var ctx = context.Background()

func RunPostgresContainer(t *testing.T) {
	pgLock.Lock()
	defer pgLock.Unlock()
	pgWg.Add(1)

	if !postgresRunning {
		runPostgresContainer(t)
		postgresRunning = true
	}
}

func runPostgresContainer(t *testing.T) testcontainers.Container {
	passwd := "passwd"
	containerRequest := testcontainers.ContainerRequest{
		Image:        "postgres:13.7",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForLog("UTC [1] LOG:  database system is ready to accept connections"),
		Env: map[string]string{
			"POSTGRES_PASSWORD": passwd,
		},
	}
	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: containerRequest,
		Started:          true,
	})
	assert.NoError(t, err)
	host, err := postgresContainer.Host(ctx)
	assert.NoError(t, err)
	port, err := postgresContainer.MappedPort(ctx, "5432")
	assert.NoError(t, err)
	logrus.WithField("host", host).
		WithField("port", port).
		Info("Postgres is running")
	viper.Set("postgres.host", host)
	viper.Set("postgres.port", port.Port())
	viper.Set("postgres.password", passwd)
	return postgresContainer
}
