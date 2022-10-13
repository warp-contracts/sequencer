package sequencerdb

import (
	"github.com/stretchr/testify/assert"
	"github.com/warp-contracts/sequencer/_tests/_testcontainers"
	"github.com/warp-contracts/sequencer/config"
	"github.com/warp-contracts/sequencer/db/conn"
	"testing"
)

func TestSequenced(t *testing.T) {
	config.Init()
	done := _testcontainers.RunPostgresContainer(t)
	defer done()

	t.Parallel()

	connection := conn.GetConnection()
	err := connection.AutoMigrate(Sequencer{})
	assert.NoError(t, err)

	t.Run("save", func(t *testing.T) {
		defer connection.Where("1=1").Delete(Sequencer{})
		origin := &Sequencer{
			OriginalSig:           "OriginalSig:          ",
			OriginalOwner:         "OriginalOwner:        ",
			OriginalAddress:       "OriginalAddress:      ",
			SequenceBlockId:       "SequenceBlockId:      ",
			SequenceBlockHeight:   123,
			SequenceTransactionId: "SequenceTransactionId:",
			SequenceMillis:        321,
			SequenceSortKey:       "SequenceSortKey:      ",
			BundlerTxId:           "BundlerTxId:          ",
			BundlerResponse:       "BundlerResponse:      ",
		}
		Save(origin)
		var saved *Sequencer
		connection.First(&saved)
		assert.Equal(t, origin, saved)
	})
}
