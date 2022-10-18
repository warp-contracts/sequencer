package routes

import (
	"bytes"
	_ "encoding/base64"
	"encoding/json"
	"github.com/everFinance/goar/types"
	_ "github.com/everFinance/goar/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/warp-contracts/sequencer/_tests/_testcontainers"
	"github.com/warp-contracts/sequencer/ar"
	"github.com/warp-contracts/sequencer/config"
	"github.com/warp-contracts/sequencer/db/conn"
	"github.com/warp-contracts/sequencer/db/interactiondb"
	"github.com/warp-contracts/sequencer/db/sequencerdb"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterSequence(t *testing.T) {
	t.Parallel()
	_testcontainers.RunPostgresContainer(t)
	initTest(t)

	c, writer := GetTestGinContext()
	c.Request.Method = http.MethodPost
	transaction := getTransactionSample()
	jsonTransaction, err := json.Marshal(transaction)
	assert.NoError(t, err)

	c.Request.Body = io.NopCloser(bytes.NewReader(jsonTransaction))

	RegisterSequencer(c)

	var bundlrResp types.BundlrResp
	assert.NoError(t, json.Unmarshal(writer.Body.Bytes(), &bundlrResp))
	t.Run("response", func(t *testing.T) {
		t.Parallel()
		t.Run("should have success status code", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, 200, writer.Code)
		})

		t.Run("id should not be empty", func(t *testing.T) {
			t.Parallel()
			assert.NotEmpty(t, bundlrResp.Id)
		})

		t.Run("should see transaction in the arweave.net", func(t *testing.T) {
			t.Parallel()
			resp, err := http.Get("https://arweave.net/" + bundlrResp.Id)
			assert.NoError(t, err)
			all, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)
			var bundlerTransaction *types.Transaction
			assert.NoError(t, json.Unmarshal(all, &bundlerTransaction))
			assert.Equal(t, transaction, bundlerTransaction)
		})
	})

	t.Run("should save data in the database", func(t *testing.T) {
		t.Parallel()
		db := conn.GetConnection()
		t.Run("interactions", func(t *testing.T) {
			t.Parallel()
			var interaction interactiondb.Interaction
			db.
				Table("interactions").
				//Where("interaction_id = ?", transaction.ID).
				First(&interaction)
			assert.NoError(t, db.Error)

			t.Run("should save interaction data", func(t *testing.T) {
				assert.NotEqual(t, interaction, interactiondb.Interaction{})
			})
			t.Run("should contain correct field values", func(t *testing.T) {
				assert.True(t, interaction.ConfirmingPeer == "https://node.bundlr.network" ||
					interaction.ConfirmingPeer == "https://node2.bundlr.network")
			})
		})
	})
}

func initTest(t *testing.T) {
	config.Init()

	connection := conn.GetConnection()
	assert.NoError(t, connection.AutoMigrate(sequencerdb.Sequence{}))
	assert.NoError(t, connection.AutoMigrate(interactiondb.Interaction{}))

	ar.StartCacheRead()
}

func getTransactionSample() *types.Transaction {
	return &types.Transaction{
		Format: 2,
		ID:     "CeYcn1VTOgHgLNspQYvrlWTLLlDggnog-negpy7pjj0",
		LastTx: "p7vc1iSP6bvH_fCeUFa9LqoV5qiyW-jdEKouAT0XMoSwrNraB9mgpi29Q10waEpO",
		Owner:  "6oyWpkM7Hk3W8e3LTpOzg6Yt8EblCjwu1xpOtKkKdMedZ8hF0X1rvUpxLC5wO5m8PHjjBiC1TwvN8kMWHN3S0DHIXg9NNhwSllOh7dBb3mj05NLan-Pc2lDNQKMWDDB4D_XamWfK4lg9LMskTQ4ZmFdqM3YoiV-uJ-e9k2SoSUV0kbdPINxwJBQRVHcDfH3yOGBSU2ZPfN-nZGfl78hbN2AxAOg2_4A_Jy1ksJmIKHg6W-nWA-mDSHbXlSDu85xnE2qDp1CnZG6jSnQhCldZf3ZoIH1AdSINdcHBW8Jk3QWKfO4pZId3AaCEFLwW8Kegt9g7bCRJO9VV7s3BVfJnv45KZ4FCC4jscYWsHWRTRSRJ2NrAXrlN8ScbjlaALPdpQFszMjhPfQZYWPhy8V0iWgwZfF4qpBFkO35FVvnp_nvJmGPjpWJVkyZESlka4zirPC8Tn6uhfhI6Rnk6Z6H9bJ-XzISxL1KUTWRQf51JoiQoHu6LtN9P0tVwPhT0Ls9xI4Zh9veAQjy98wLySGqxdMrdBDZO4dwCHTOojROxaxPLW9rmroUBtgufYQui_tgJ8QAI2EOEXIU0dO8bRNCoDgTu9X8E6f0cuY4ugNR1-n0-eV7CahkcmzXACLm8i8Dm33ZTsxf5sI-fsTAGNk4HtvYgXm5wRG4VJj_M0o5nOHs",
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
				Value: "V3M5aGhZY2tjLXpTblZtYkJlcDZxX2taRDV6bXpZekRtZ01DNTBuTWl1RQ",
			},
			{
				Name:  "SW5wdXQ",
				Value: "eyJmdW5jdGlvbiI6IndoYXRldmVyIn0",
			},
		},
		Target:    "",
		Quantity:  "0",
		Data:      "NDExNQ",
		DataSize:  "4",
		DataRoot:  "CaMcorJHEzSadSFB6IfHrDQzeuBNtybdksyXC_efm6Q",
		Reward:    "72600854",
		Signature: "jiLdLnAWkhVklU6FUCpjjWO69koRxNH4at7VzWBpETzmOIFSmAFJw06UHqyxtL4Xkzcj4wyWUDym9UHZxRVyqTBHBIcJMSDpZTSO2hfIKWOlBu0aRjGZIHJDUzZ5QWxNo6pGuSDAOqLyI1tu891vegZTYiUfNUIm154dbLzSorzRQrFgmrcUIHvQzcZ5jzU28WKJ90vAIrJyrKK3YeJ4u5ZPZaHQc9UQH1jwhIFcNfvUst4XTo_1CEawnRIMLPyfAaPOmNpVJwvYzHrc-rA0ixgOlWauv1sgmKtNOr6hn1ERItiM8lRFDzK4nWCM6UMqmFO74yco6RjzFUzVQk8tIUUaM_hrEdkap1OtKcebxV4KYSkSHhAIgHMjnN6vRb4t57hUZ0yEYhpv_dYrpkq3FwzMazL5kAF4HzE8PMflOmIPHkr670ogMEYQBgG008aJblkLqNlW1LR0dPa0psMX8xmwOn5zdNRSHAfolRhje72k21AkEpb0J-yELpcN1Yqzps3JAHeoufu0sjkL-aBQ8qPc0JeJ1Ua89Ld1iprR97-N1er8wW1Z10H5wj30xTwygJYdJgj9YS8v1xsgNPxXYHS3N-zXZRfg4AcgGRSX26PKValSNJiMYFWwmU6Lz-6Q3i3tAYCR9wTc9SQ9QJJemB1mQTBiUeCwq26ANAOcDHE",
		Chunks:    nil,
	}
}
func GetTestGinContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx, w
}
