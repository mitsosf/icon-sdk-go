package iconsdk

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestIRC2(t *testing.T) {
	t.Run("Get name", func(t *testing.T) {
		url := "https://lisbon.net.solidwallet.io/api/v3"
		iconService := NewIconService(&url)
		irc2 := NewIRC2("cx273548dff8bb77ffaac5a342c4c04aeae0bc48fa", *iconService)
		res, err := irc2.Name()
		if err != nil {
			fmt.Println(err)
			assert.NotNil(t, res)
			return
		}
		fmt.Println(res)
		assert.Equal(t, "2.0", res["jsonrpc"])
		assert.Equal(t, "MyIRC2Token", res["result"])
	})

	t.Run("Get symbol", func(t *testing.T) {
		url := "https://lisbon.net.solidwallet.io/api/v3"
		iconService := NewIconService(&url)
		irc2 := NewIRC2("cx273548dff8bb77ffaac5a342c4c04aeae0bc48fa", *iconService)
		res, err := irc2.Symbol()
		if err != nil {
			fmt.Println(err)
			assert.NotNil(t, res)
			return
		}
		fmt.Println(res)
		assert.Equal(t, "2.0", res["jsonrpc"])
		assert.Equal(t, "MIT", res["result"])
	})

	t.Run("Get decimals", func(t *testing.T) {
		url := "https://lisbon.net.solidwallet.io/api/v3"
		iconService := NewIconService(&url)
		irc2 := NewIRC2("cx273548dff8bb77ffaac5a342c4c04aeae0bc48fa", *iconService)
		res, err := irc2.Decimals()
		if err != nil {
			fmt.Println(err)
			assert.NotNil(t, res)
			return
		}
		fmt.Println(res)
		assert.Equal(t, "2.0", res["jsonrpc"])
		assert.Equal(t, "0x12", res["result"])
	})

	t.Run("Get total supply", func(t *testing.T) {
		url := "https://lisbon.net.solidwallet.io/api/v3"
		iconService := NewIconService(&url)
		irc2 := NewIRC2("cx273548dff8bb77ffaac5a342c4c04aeae0bc48fa", *iconService)
		res, err := irc2.TotalSupply()
		if err != nil {
			fmt.Println(err)
			assert.NotNil(t, res)
			return
		}
		fmt.Println(res)
		assert.Equal(t, "2.0", res["jsonrpc"])
		assert.Equal(t, "0x3635c9adc5dea00000", res["result"])
	})

	t.Run("Get balance of", func(t *testing.T) {
		url := "https://lisbon.net.solidwallet.io/api/v3"
		iconService := NewIconService(&url)
		irc2 := NewIRC2("cx273548dff8bb77ffaac5a342c4c04aeae0bc48fa", *iconService)
		res, err := irc2.BalanceOf("hx8dc6ae3d93e60a2dddf80bfc5fb1cd16a2bf6160")
		if err != nil {
			fmt.Println(err)
			assert.NotNil(t, res)
			return
		}
		fmt.Println(res)
		assert.Equal(t, "2.0", res["jsonrpc"])
		assert.Nil(t, res["error"])
		assert.NotNil(t, res["result"])
	})

	t.Run("Transfer irc2", func(t *testing.T) {
		url := "https://lisbon.net.solidwallet.io/api/v3"
		iconService := NewIconService(&url)
		irc2 := NewIRC2("cx273548dff8bb77ffaac5a342c4c04aeae0bc48fa", *iconService)
		privateKey := "3468ea815d8896ef4552f10768caf2660689b965975c3ec2c1f5fe84bc3a77a5"
		wallet := NewWallet(&privateKey)
		res, err := irc2.Transfer(*wallet, "hx8dc6ae3d93e60a2dddf80bfc5fb1cd16a2bf6160", "12.317", "0x3", "0x2", "0x1", "0x186a00")
		if err != nil {
			fmt.Println(err)
			assert.NotNil(t, res)
			return
		}
		fmt.Println(res)
		assert.Equal(t, "2.0", res["jsonrpc"])
		assert.Nil(t, res["error"])
		assert.NotNil(t, res["result"])
		assert.True(t, strings.HasPrefix(res["result"].(string), "0x"), "result does not start with '0x'")
	})

}
