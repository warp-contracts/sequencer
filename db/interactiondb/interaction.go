package interactiondb

import (
	"github.com/warp-contracts/sequencer/db/conn"
)

func Save(s *Interaction) error {
	return conn.GetConnection().Create(s).Error
}

type Interaction struct {
	InteractionId      string
	Interaction        string
	BlockHeight        int64
	BlockId            string
	ContractId         string
	Function           string
	Input              string
	ConfirmationStatus string
	ConfirmingPeer     string
	Source             string
	BundlerTxId        string
	InteractWrite      string
	SortKey            string
	Evolve             string
}
