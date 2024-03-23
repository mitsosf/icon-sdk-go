package iconsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

type Transaction struct {
	iconService *IconService
	data        map[string]interface{}
}

func NewTransaction(iconService *IconService) *Transaction {
	return &Transaction{
		data: map[string]interface{}{
			"jsonrpc": "2.0",
			"method":  "",
			"id":      rand.Intn(9999),
		},
		iconService: iconService,
	}
}

func (t *Transaction) SetMethod(method string) {
	t.data["method"] = method
}

func (t *Transaction) SetParams(newParams map[string]interface{}) {
	params, ok := t.data["params"].(map[string]interface{})
	if !ok {
		params = make(map[string]interface{})
		t.data["params"] = params
	}

	for key, value := range newParams {
		params[key] = value
	}
}

func (t *Transaction) Send() (map[string]interface{}, error) {
	client := &http.Client{}

	jsonData, err := json.Marshal(t.data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request data: %w", err)
	}

	req, err := http.NewRequest("POST", t.iconService.IconServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("creating request failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyText, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected response status: %s, body: %s", resp.Status, string(bodyText))
	}

	var data map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	return data, nil
}
