package sequencerdb

import (
	"github.com/warp-contracts/sequencer/db/conn"
)

func Save(s *Sequence) error {
	return conn.GetConnection().Create(s).Error
}

type Sequence struct {
	OriginalSig           string
	OriginalOwner         string
	OriginalAddress       string
	SequenceBlockId       string
	SequenceBlockHeight   int64
	SequenceTransactionId string
	SequenceMillis        string
	SequenceSortKey       string
	BundlerTxId           string
	BundlerResponse       string
}

func (*Sequence) TableName() string {
	return "sequencer"
}
