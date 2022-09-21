package ar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetArweaveClient(t *testing.T) {
	info, err := GetArweaveClient().GetInfo()
	assert.NoError(t, err)
	assert.NotNil(t, info)
}
