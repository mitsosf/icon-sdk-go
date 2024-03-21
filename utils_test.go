package iconsdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtils(t *testing.T) {
	t.Run("Hex to ICX", func(t *testing.T) {
		expected := "1.31231232"
		result := hexToIcx("0x1236451a23f80000", nil)
		assert.Equal(t, expected, result)
	})

	t.Run("ICX to hex", func(t *testing.T) {
		expected := "0x1236451a23f80000"
		result := icxToHex("1.31231232", nil)
		assert.Equal(t, expected, result)
	})

	t.Run("ICX to hex trailing 00", func(t *testing.T) {
		expected := "0x1236451a23f80000"
		result := icxToHex("1.312312320000000", nil)
		assert.Equal(t, expected, result)
	})
}
