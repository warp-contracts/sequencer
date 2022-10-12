package sequencerdb

import (
	"github.com/stretchr/testify/assert"
	"github.com/warp-contracts/sequencer/_tests/_testcontainers"
	"github.com/warp-contracts/sequencer/config"
	"testing"
)

func TestSequenced(t *testing.T) {
	config.Init("../..")
	done := _testcontainers.RunPostgresContainer(t)
	defer done()

	t.Parallel()

	connection := getConnection()
	err := connection.AutoMigrate(Sequence{})
	assert.NoError(t, err)

	t.Run("save", func(t *testing.T) {
		defer connection.Where("1=1").Delete(Sequence{})
		origin := &Sequence{
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
		var saved *Sequence
		connection.First(&saved)
		assert.Equal(t, origin, saved)
	})
}
