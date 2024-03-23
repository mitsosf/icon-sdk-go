package iconsdk

import (
	"fmt"
	"math/big"
	"strings"
)

func icxToHex(value string, decimals *int) string {
	if decimals == nil {
		decimals = new(int)
		*decimals = 18
	}
	val, _, err := big.ParseFloat(value, 10, 0, big.ToNearestEven)
	if err != nil {
		fmt.Printf("Error parsing value: %v\n", err)
		return ""
	}

	multiplier := new(big.Float).SetFloat64(float64(1))
	for i := 0; i < *decimals; i++ {
		multiplier.Mul(multiplier, big.NewFloat(10))
	}
	val.Mul(val, multiplier)

	valInt, _ := val.Int(nil) // This truncates the decimal part

	hexStr := "0x" + valInt.Text(16)

	return hexStr
}

func hexToIcx(value string, decimals *int) string {
	if decimals == nil {
		decimals = new(int)
		*decimals = 18
	}

	if strings.HasPrefix(value, "0x") {
		value = value[2:]
	}

	valInt, success := new(big.Int).SetString(value, 16)
	if !success {
		fmt.Println("Error parsing hex value")
		return ""
	}

	val := new(big.Float).SetInt(valInt)

	divisor := new(big.Float).SetFloat64(float64(1))
	for i := 0; i < *decimals; i++ {
		divisor.Mul(divisor, big.NewFloat(10))
	}

	result := new(big.Float).Quo(val, divisor)

	resultStr := fmt.Sprintf("%.*f", *decimals, result)
	return trimTrailingZeros(resultStr)
}

func trimTrailingZeros(s string) string {
	s = strings.TrimRight(s, "0")
	// Ensure there's a digit after the decimal point
	if s[len(s)-1] == '.' {
		s += "0"
	}
	return s
}
