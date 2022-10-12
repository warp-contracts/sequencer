package sequencerdb

import (
	"github.com/warp-contracts/sequencer/db/conn"
	"gorm.io/gorm"
)

const tableName = "sequencer"

func Save(s *Sequence) {
	getConnection().Create(s)
}

func getConnection() *gorm.DB {
	connection := conn.GetConnection()
	connection.Table(tableName)
	return connection
}

type Sequence struct {
	OriginalSig           string
	OriginalOwner         string
	OriginalAddress       string
	SequenceBlockId       string
	SequenceBlockHeight   int64
	SequenceTransactionId string
	SequenceMillis        int64
	SequenceSortKey       string
	BundlerTxId           string
	BundlerResponse       string
}
