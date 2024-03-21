package iconsdk

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/sha3"
	"sort"
	"strings"
)

// serializeTransaction serializes JSON data with custom formatting and optionally hashes it.
func serializeTransaction(data interface{}, hashed bool) (string, error) {
	// Marshal the data into JSON bytes
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Unmarshal into a generic interface for manual traversal
	var jsonData interface{}
	if err := json.Unmarshal(jsonBytes, &jsonData); err != nil {
		return "", err
	}

	result := "icx_sendTransaction." + valueTraverse(jsonData, true)

	if hashed {
		// If hashed is true, hash the result string
		hash := sha3.New256()
		hash.Write([]byte(result))
		return hex.EncodeToString(hash.Sum(nil)), nil
	}
	return result, nil
}

// valueTraverse traverses the JSON data recursively, applying custom formatting.
func valueTraverse(value interface{}, external bool) string {
	switch v := value.(type) {
	case map[string]interface{}:
		var keys []string
		for k := range v {
			keys = append(keys, k)
		}
		sort.Strings(keys) // Sort the keys for consistent ordering

		var result bytes.Buffer
		if !external {
			result.WriteString("{")
		}
		for _, k := range keys {
			result.WriteString(fmt.Sprintf("%s.", k))
			result.WriteString(valueTraverse(v[k], false))
			result.WriteString(".")
		}
		trimTrailingDot(&result)
		if !external {
			result.WriteString("}")
		}
		return result.String()

	case []interface{}:
		var result bytes.Buffer
		result.WriteString("[")
		for _, elem := range v {
			result.WriteString(valueTraverse(elem, false))
			result.WriteString(".")
		}
		trimTrailingDot(&result)
		result.WriteString("]")
		return result.String()

	case string:
		return escapeString(v)

	case float64:
		return fmt.Sprintf("%v", v)

	case bool:
		return fmt.Sprintf("%t", v)

	case nil:
		return "\\0"

	default:
		return ""
	}
}

// trimTrailingDot removes a trailing dot from the buffer.
func trimTrailingDot(buf *bytes.Buffer) {
	if buf.Len() > 0 && buf.Bytes()[buf.Len()-1] == '.' {
		buf.Truncate(buf.Len() - 1)
	}
}

// escapeString applies custom escaping rules to a string.
func escapeString(value string) string {
	replacer := strings.NewReplacer("\\", "\\\\", ".", "\\.", "{", "\\{", "}", "\\}", "[", "\\[", "]", "\\]")
	return replacer.Replace(value)
}
