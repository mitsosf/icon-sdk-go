package iconsdk

// IconService holds the URL to the ICON service.
type IconService struct {
	IconServiceURL string
}

// NewIconService creates a new instance of IconService with a default or specified URL.
func NewIconService(iconServiceURL *string) *IconService {
	defaultURL := "https://api.icon.community/api/v3"
	if iconServiceURL == nil {
		return &IconService{IconServiceURL: defaultURL}
	}
	return &IconService{IconServiceURL: *iconServiceURL}
}

func (i *IconService) GetLastBlock() (map[string]interface{}, error) {
	transaction := NewTransactionBuilder(i).
		Method("icx_getLastBlock").
		Build()
	return transaction.Send()
}

func (i *IconService) GetBlockByHeight(height string) (map[string]interface{}, error) {
	transaction := NewTransactionBuilder(i).
		Method("icx_getBlockByHeight").
		BlockHeight(height).
		Build()
	return transaction.Send()
}

func (i *IconService) GetBlockByHash(hash string) (map[string]interface{}, error) {
	transaction := NewTransactionBuilder(i).
		Method("icx_getBlockByHash").
		BlockHash(hash).
		Build()
	return transaction.Send()
}

func (i *IconService) GetBalance(address string) (map[string]interface{}, error) {
	transaction := NewTransactionBuilder(i).
		Method("icx_getBalance").
		Address(address).
		Build()
	return transaction.Send()
}

func (i *IconService) GetTransactionResult(txHash string) (map[string]interface{}, error) {
	transaction := NewTransactionBuilder(i).
		Method("icx_getTransactionResult").
		TxHash(txHash).
		Build()
	return transaction.Send()
}

func (i *IconService) Call(score string, params map[string]interface{}) (map[string]interface{}, error) {
	transaction := NewTransactionBuilder(i).
		Method("icx_call").
		To(score).
		Call(params).
		Build()
	return transaction.Send()
}

func (i *IconService) SendTransaction(privateKey string, from string, to string, value string, version string, nid string, nonce string, stepLimit string) (map[string]interface{}, error) {
	transaction := NewTransactionBuilder(i).
		Method("icx_sendTransaction").
		From(from).
		To(to).
		Value(value).
		Version(version).
		Nid(nid).
		Timestamp().
		Nonce(nonce).
		StepLimit(stepLimit).
		Sign(privateKey).
		Build()
	return transaction.Send()
}

func (i *IconService) SendTransactionWithMessage(privateKey string, from string, to string, value string, version string, nid string, nonce string, stepLimit string, message string) (map[string]interface{}, error) {
	transaction := NewTransactionBuilder(i).
		Method("icx_sendTransaction").
		From(from).
		To(to).
		Value(value).
		Version(version).
		Nid(nid).
		Timestamp().
		Nonce(nonce).
		StepLimit(stepLimit).
		Message(message).
		Sign(privateKey).
		Build()
	return transaction.Send()
}
