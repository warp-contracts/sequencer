package interactiondb

import (
	"github.com/warp-contracts/sequencer/db/conn"
	"gorm.io/gorm"
)

func Save(s *Interaction) {
	getConnection().Create(s)
}

func getConnection() *gorm.DB {
	connection := conn.GetConnection()
	return connection
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
