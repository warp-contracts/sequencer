package eth

import (
	"github.com/everFinance/goar/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	t.Skip("Not implemented yet")
	t.Parallel()

	t.Run("should return true when transaction is signed correctly", func(t *testing.T) {
		tx := types.Transaction{
			Format: 2,
			ID:     "NoUZcl1NoYwcngUJ7nZcDDqufXFeADVyzFgV0KjgWZw",
			LastTx: "p7vc1iSP6bvH_fCeUFa9LqoV5qiyW-jdEKouAT0XMoSwrNraB9mgpi29Q10waEpO",
			Owner:  "0x3a709448077e6ef6beda8f664a36cc24abb39002",
			Tags: []types.Tag{
				{
					Name:  "QXBwLU5hbWU",
					Value: "U21hcnRXZWF2ZUFjdGlvbg",
				},
				{
					Name:  "QXBwLVZlcnNpb24",
					Value: "MC4zLjA",
				},
				{
					Name:  "U0RL",
					Value: "V2FycA",
				},
				{
					Name:  "Q29udHJhY3Q",
					Value: "NDhHX0lsbFU5Ry1QUnlsNE9kczg4U1R0UTFoMEVvOHpIUVVIZE5sSEtadw",
				},
				{
					Name:  "SW5wdXQ",
					Value: "eyJmdW5jdGlvbiI6InBvc3RNZXNzYWdlIiwiY29udGVudCI6InRlc3Q1In0",
				},
				{
					Name:  "U2lnbmF0dXJlLVR5cGU",
					Value: "ZXRoZXJldW0",
				},
			},
			Target:    "",
			Quantity:  "0",
			Data:      "OTE5Mg",
			DataSize:  "4",
			DataRoot:  "qPnASO7WoQzjrbxVN4YFp5hskttCPVSTt3M8gXOtPcg",
			Reward:    "72600854",
			Signature: "0x551b0d367c29c67e06a1cf46d023aa595e13e21b0b73e7096015dd8627fcb06346075ceb135e342320c7040c3ce327547d9a158400d14c7f51acda82da3cfa661c",
		}
		assert.True(t, Validate(tx))
	})

	t.Run("should return false when transaction is not signed", func(t *testing.T) {
		assert.False(t, Validate(types.Transaction{}))
	})
}
