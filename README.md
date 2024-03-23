<p align="center">
  <img 
    src="https://iconation.team/images/very_small.png" 
    width="120px"
    alt="ICONation logo">
</p>

<h1 align="center">ICON SDK for Go</h1>

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This is an SDK to communicate with the ICON blockchain, built for Go.

Disclaimer: I cannot guarantee optimal performance of this software. It is provided as is and without any assurances. Use it at your own risk.

Features
--------

- Wallet management
- Read data from the blockchain
- Send ICX transactions
- Perform SCORE calls
- Transaction builder


Testing
--------
To run tests, ensure you have Go installed and run:
```shell
go test
```

Usage
--------
```go
import (
    "fmt"
    "github.com/mitsosf/icon-sdk-go"
)

iconservice := iconsdk.NewIconService(nil)
res, err := iconservice.GetBalance("<address>")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(res)
	
// In case you want to initalize IconService on a testnet
iconServiceUrl := "https://lisbon.net.solidwallet.io/api/v3"
iconservice := iconsdk.NewIconService(&iconServiceUrl)

// Creating or loading a wallet
wallet := iconsdk.NewWallet(nil)
privateKey := "01234..."
wallet := iconsdk.NewWallet(&privateKey)

// IRC2
irc2 = iconsdk.NewIRC2("cx123...", *iconservice)
```
