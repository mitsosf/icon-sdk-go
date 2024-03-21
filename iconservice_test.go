package iconsdk

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIconService(t *testing.T) {
	t.Run("Get last block", func(t *testing.T) {
		iconService := NewIconService(nil)
		block, err := iconService.GetLastBlock()
		if err != nil {
			fmt.Println(err)
			assert.NotNil(t, block)
			return
		}
		fmt.Println(block)
		assert.Equal(t, "2.0", block["jsonrpc"])
	})

	t.Run("Get block by height", func(t *testing.T) {
		iconService := NewIconService(nil)
		block, err := iconService.GetBlockByHeight("0x0")
		if err != nil {
			fmt.Println(err)
			assert.NotNil(t, block)
			return
		}
		fmt.Println(block)
		assert.Equal(t, "2.0", block["jsonrpc"])
	})

	t.Run("Get block by hash", func(t *testing.T) {
		iconService := NewIconService(nil)
		block, err := iconService.GetBlockByHash("0xcf43b3fd45981431a0e64f79d07bfcf703e064b73b802c5f32834eec72142190")
		if err != nil {
			fmt.Println(err)
			assert.NotNil(t, block)
			return
		}
		fmt.Println(block)
		assert.Equal(t, "2.0", block["jsonrpc"])
	})

	t.Run("Get balance", func(t *testing.T) {
		iconService := NewIconService(nil)
		block, err := iconService.GetBalance("hxd5ace539bf910635c2fa0e9c185d2d3c8d52c4cc")
		if err != nil {
			fmt.Println(err)
			assert.NotNil(t, block)
			return
		}
		fmt.Println(block)
		assert.Equal(t, "2.0", block["jsonrpc"])
	})

	t.Run("Get transaction result", func(t *testing.T) {
		iconService := NewIconService(nil)
		block, err := iconService.GetTransactionResult("0x1b6133792cee1ab2e54ae68faf9f49daf81c7e46d68b1ca281acc718602c77dd")
		if err != nil {
			fmt.Println(err)
			assert.NotNil(t, block)
			return
		}
		fmt.Println(block)
		assert.Equal(t, "2.0", block["jsonrpc"])
	})

	t.Run("Call", func(t *testing.T) {
		iconService := NewIconService(nil)
		params := make(map[string]interface{})
		params["method"] = "balanceOf"
		params["params"] = map[string]interface{}{"_owner": "hx70e8eeb5d23ab18a828ec95f769db6d953e5f0fd"}
		block, err := iconService.Call("cx9ab3078e72c8d9017194d17b34b1a47b661945ca", params)
		if err != nil {
			fmt.Println(err)
			assert.NotNil(t, block)
			return
		}
		fmt.Println(block)
		assert.Equal(t, "2.0", block["jsonrpc"])
	})

	t.Run("Send transaction", func(t *testing.T) {
		url := "https://lisbon.net.solidwallet.io/api/v3"
		iconService := NewIconService(&url)

		result, err := iconService.SendTransaction(
			"f4ade1ff528c9e0bf10d35909e3486ef6ce88df8a183fc1cc2c65bfa9a53d3fd",
			"hxb14e0c751899676a1a4e655a34063b42260f844b",
			"hxf8689d6c4c8f333651469fdea2ac59a18f6c2421",
			"11.291182",
			"0x3",
			"0x2",
			"0x1",
			"0x186a0",
		)
		if err != nil {
			fmt.Println(err)
			assert.NotNil(t, result)
			return
		}
		fmt.Println(result)
		assert.Equal(t, "2.0", result["jsonrpc"])
	})

	t.Run("Send transaction with hex ICX value", func(t *testing.T) {
		url := "https://lisbon.net.solidwallet.io/api/v3"
		iconService := NewIconService(&url)

		result, err := iconService.SendTransaction(
			"f4ade1ff528c9e0bf10d35909e3486ef6ce88df8a183fc1cc2c65bfa9a53d3fd",
			"hxb14e0c751899676a1a4e655a34063b42260f844b",
			"hxf8689d6c4c8f333651469fdea2ac59a18f6c2421",
			"0x1236451a23f80000",
			"0x3",
			"0x2",
			"0x1",
			"0x186a0",
		)
		if err != nil {
			fmt.Println(err)
			assert.NotNil(t, result)
			return
		}
		fmt.Println(result)
		assert.Equal(t, "2.0", result["jsonrpc"])
	})

	t.Run("Send transaction with message", func(t *testing.T) {
		url := "https://lisbon.net.solidwallet.io/api/v3"
		iconService := NewIconService(&url)

		result, err := iconService.SendTransactionWithMessage(
			"f4ade1ff528c9e0bf10d35909e3486ef6ce88df8a183fc1cc2c65bfa9a53d3fd",
			"hxb14e0c751899676a1a4e655a34063b42260f844b",
			"hxf8689d6c4c8f333651469fdea2ac59a18f6c2421",
			"0x1236451a23f80000",
			"0x3",
			"0x2",
			"0x1",
			"0x186a00",
			"Test message",
		)
		if err != nil {
			fmt.Println(err)
			assert.NotNil(t, result)
			return
		}
		fmt.Println(result)
		assert.Equal(t, "2.0", result["jsonrpc"])
	})
}
