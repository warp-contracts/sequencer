package interactiondb

import (
	"github.com/stretchr/testify/assert"
	"github.com/warp-contracts/sequencer/_tests/_testcontainers"
	"github.com/warp-contracts/sequencer/config"
	"testing"
)

func TestInteraction(t *testing.T) {
	config.Init()
	done := _testcontainers.RunPostgresContainer(t)
	defer done()

	t.Parallel()

	connection := getConnection()
	err := connection.AutoMigrate(Interaction{})
	assert.NoError(t, err)

	t.Run("save", func(t *testing.T) {
		defer connection.Where("1=1").Delete(Interaction{})
		origin := &Interaction{}
		Save(origin)
		var saved *Interaction
		connection.First(&saved)
		assert.Equal(t, origin, saved)
	})
}
