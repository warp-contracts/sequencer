package conn

import (
	"github.com/stretchr/testify/assert"
	"github.com/warp-contracts/sequencer/_tests/_testcontainers"
	"github.com/warp-contracts/sequencer/config"
	"testing"
)

func TestDbConnection(t *testing.T) {
	t.Parallel()

	config.Init()
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
