package utils

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strconv"
	"strings"
)

// IntSliceToString converts slice of integers to a string
func IntSliceToString(values []int) string {
	valuesText := []string{}
	result := ""
	for i := range values {
		text := strconv.Itoa(values[i])
		valuesText = append(valuesText, text)
		result = strings.Join(valuesText, "")
	}

	return result
}

// GenerateNumericCode generates numeric code based on length
func GenerateNumericCode(length int) (string, error) {
	if length <= 0 {
		// Depending on your requirements, you could return an error or a specific behavior for non-positive lengths
		return "", errors.New("length must be larger than 0")
	}

	var builder strings.Builder
	for i := 0; i < length; i++ {
		// Use crypto/rand for secure randomness; it's more appropriate for sensitive codes
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			// Handle the potential error from the crypto/rand usage
			return "", err
		}
		builder.WriteString(num.String())
	}

	return builder.String(), nil
}
