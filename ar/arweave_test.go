package ar

import (
	"github.com/stretchr/testify/assert"
	"github.com/warp-contracts/sequencer/config"
	"testing"
)

func TestGetArweaveClient(t *testing.T) {
	config.Init()
	client := GetArweaveClient()
	assert.NotNil(t, client)
	t.Run("should return info", func(t *testing.T) {
		info, err := client.GetInfo()
		assert.NoError(t, err)
		assert.NotEmpty(t, info)
	})
}
