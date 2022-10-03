package db

import (
	"github.com/stretchr/testify/assert"
	"github.com/warp-contracts/gateway/_tests/_testcontainers"
	"github.com/warp-contracts/gateway/config"
	"testing"
)

func TestDbConnection(t *testing.T) {
	config.Init()
	//ctx := context.Background()
	//port, err := freeport.GetFreePort()
	//passwd := "passwd"
	//containerRequest := testcontainers.ContainerRequest{
	//	Image: "postgres:13.7",
	//	ExposedPorts: []string{fmt.Sprintf("%d/tcp", port)},
	//	WaitingFor:   wait.ForLog("database system is ready to accept connections"),
	//	Env: map[string]string{
	//		"POSTGRES_PASSWORD": passwd,
	//	},Ž
	//}
	//postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
	//	ContainerRequest: containerRequest,
	//	Started:          true,
	//})
	//assert.NoError(t, err)
	//defer postgresContainer.Terminate(ctx)
	done := _testcontainers.RunPostgresContainer(t)
	defer done()

	t.Run("should connect to DB", func(t *testing.T) {
		conn := GetPostgresConnection()
		assert.NotNil(t, conn)

		rows, err := conn.Query("SELECT 1")
		assert.NoError(t, err)
		assert.NotNil(t, rows)
	})
}
