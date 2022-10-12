package sequencerdb

import (
	"github.com/warp-contracts/sequencer/db/conn"
)

const tableName = "sequencer"

func Save(s *Sequence) {
	connection := conn.GetConnection()
	connection.Table(tableName)
	connection.Create(s)
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
