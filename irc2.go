package iconsdk

type IRC2 struct {
	iconService IconService
	contract    string
}

func NewIRC2(contract string, iconService IconService) *IRC2 {
	return &IRC2{contract: contract, iconService: iconService}
}

func (i *IRC2) Name() (map[string]interface{}, error) {
	params := make(map[string]interface{})
	params["method"] = "name"

	transaction := NewTransactionBuilder(&i.iconService).
		Method("icx_call").
		To(i.contract).
		Call(params).
		Build()

	return transaction.Send()
}

func (i *IRC2) Symbol() (map[string]interface{}, error) {
	params := make(map[string]interface{})
	params["method"] = "symbol"

	transaction := NewTransactionBuilder(&i.iconService).
		Method("icx_call").
		To(i.contract).
		Call(params).
		Build()

	return transaction.Send()
}

func (i *IRC2) Decimals() (map[string]interface{}, error) {
	params := make(map[string]interface{})
	params["method"] = "decimals"

	transaction := NewTransactionBuilder(&i.iconService).
		Method("icx_call").
		To(i.contract).
		Call(params).
		Build()

	return transaction.Send()
}

func (i *IRC2) TotalSupply() (map[string]interface{}, error) {
	params := make(map[string]interface{})
	params["method"] = "totalSupply"

	transaction := NewTransactionBuilder(&i.iconService).
		Method("icx_call").
		To(i.contract).
		Call(params).
		Build()

	return transaction.Send()
}

func (i *IRC2) BalanceOf(address string) (map[string]interface{}, error) {
	params := make(map[string]interface{})
	params["method"] = "balanceOf"
	params["params"] = map[string]interface{}{"_owner": address}

	transaction := NewTransactionBuilder(&i.iconService).
		Method("icx_call").
		To(i.contract).
		Call(params).
		Build()

	return transaction.Send()
}

func (i *IRC2) Transfer(wallet Wallet, to string, value string, version string, nid string, nonce string, stepLimit string) (map[string]interface{}, error) {
	if value[0:2] != "0x" {
		value = icxToHex(value, nil)
	}
	params := make(map[string]interface{})
	params["method"] = "transfer"
	params["params"] = map[string]interface{}{"_to": to, "_value": value}

	transaction := NewTransactionBuilder(&i.iconService).
		Method("icx_sendTransaction").
		From(wallet.PublicAddress).
		To(i.contract).
		Version(version).
		Nid(nid).
		Timestamp().
		Nonce(nonce).
		StepLimit(stepLimit).
		Call(params).
		Sign(wallet).
		Build()

	return transaction.Send()
}
