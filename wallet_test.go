package iconsdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Create wallet", func(t *testing.T) {
		wallet := NewWallet(nil)
		assert.NotNil(t, wallet)
	})

	t.Run("Import wallet", func(t *testing.T) {
		privateKey := "f4ade1ff528c9e0bf10d35909e3486ef6ce88df8a183fc1cc2c65bfa9a53d3fd"
		wallet := NewWallet(&privateKey)
		expected := "hxb14e0c751899676a1a4e655a34063b42260f844b"
		assert.Equal(t, expected, wallet.PublicAddress)
	})
}
