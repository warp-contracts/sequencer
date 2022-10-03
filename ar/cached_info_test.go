package ar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCache(t *testing.T) {
	t.Run("should panic when request cache info ", func(t *testing.T) {
		assertPanic(t, func() {
			GetCachedInfo()
		})
	})
	StartCacheRead()

	t.Run("should return cached info", func(t *testing.T) {
		assert.NotNil(t, GetCachedInfo())
	})
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
