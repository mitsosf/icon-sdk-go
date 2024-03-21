package iconsdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerializeTransaction(t *testing.T) {
	t.Run("Serialize without hashing", func(t *testing.T) {
		data := map[string]interface{}{
			"from":  "hx123",
			"to":    "hx456",
			"value": "0x123",
		}
		expected := "icx_sendTransaction.from.hx123.to.hx456.value.0x123"
		result, err := serializeTransaction(data, false)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Serialize with hashing", func(t *testing.T) {
		data := map[string]interface{}{
			"to":    "hx456",
			"value": "0x123",
			"from":  "hx123",
		}
		expected := "9173de2fb61aabbb9f7a4bff3e7ab9c847735a296feaa98c3d34bd4a1b7ea188"
		result, err := serializeTransaction(data, true)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Complex nested JSON", func(t *testing.T) {
		data := map[string]interface{}{
			"from": "hx123",
			"data": map[string]interface{}{
				"payload": []interface{}{
					"item1",
					map[string]interface{}{
						"key": "value",
					},
				},
			},
			"to": "hx456",
		}
		expected := "icx_sendTransaction.data.{payload.[item1.{key.value}]}.from.hx123.to.hx456"
		result, err := serializeTransaction(data, false)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Edge cases", func(t *testing.T) {
		data := map[string]interface{}{
			"empty":        "",
			"specialChars": "\\.{[}]",
		}
		expected := "icx_sendTransaction.empty..specialChars.\\\\\\.\\{\\[\\}\\]"
		result, err := serializeTransaction(data, false)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}
