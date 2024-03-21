package iconsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

type Transaction struct {
	jsonrpc     string
	id          int
	iconService *IconService
	method      string
	params      map[string]interface{}
}

func NewTransaction(iconService *IconService) *Transaction {
	return &Transaction{
		jsonrpc:     "2.0",
		id:          rand.Intn(9999),
		iconService: iconService,
		params:      make(map[string]interface{}),
	}
}

// SetMethod sets the method for the transaction
func (t *Transaction) SetMethod(method string) {
	t.method = method
}

// SetParams updates or adds new params to the transaction
func (t *Transaction) SetParams(newParams map[string]interface{}) {
	for key, value := range newParams {
		t.params[key] = value
	}
}

// send sends a request to the ICON service and returns the response.
func (iconservice *IconService) Send(data interface{}) (*json.RawMessage, error) {
	client := &http.Client{}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request data: %w", err)
	}

	req, err := http.NewRequest("POST", iconservice.IconServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("creating request failed: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyText, _ := ioutil.ReadAll(resp.Body) // Ignoring error on purpose for simplicity
		return nil, fmt.Errorf("unexpected response status: %s, body: %s", resp.Status, string(bodyText))
	}

	var result json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decoding response failed: %w", err)
	}

	return &result, nil
}
