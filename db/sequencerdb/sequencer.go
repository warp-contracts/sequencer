package sequencerdb

import (
	"github.com/warp-contracts/sequencer/db/conn"
)

func Save(s *Sequencer) error {
	return conn.GetConnection().Create(s).Error
}

type Sequencer struct {
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
