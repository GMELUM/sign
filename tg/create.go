package tg

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// Create generates a data-check-string, computes the HMAC-SHA256 signature, and returns the signed parameters.
func Create(params map[string]interface{}, secret string) (string, error) {
	// Prepare a sorted slice of key-value pairs.
	var pairs []string
	for key, value := range params {
		// Convert the value to a string.
		// Assuming that all values in the map can be converted to strings directly.
		valueStr := fmt.Sprintf("%v", value)
		pairs = append(pairs, fmt.Sprintf("%s=%s", key, valueStr))
	}

	// Sort the key-value pairs lexicographically by key.
	sort.Strings(pairs)

	// Create the "data-check-string" by joining the sorted pairs with newline characters.
	dataCheckString := strings.Join(pairs, "\n")

	// Create the secret key (hash of bot_token).
	hashSecret := sha256.Sum256([]byte(secret))

	// Compute the HMAC-SHA256 of the "data-check-string" using the secret key.
	impHmac := hmac.New(sha256.New, hashSecret[:])
	impHmac.Write([]byte(dataCheckString))
	computedHash := hex.EncodeToString(impHmac.Sum(nil))

	// Add the computed hash to the parameters.
	params["hash"] = computedHash

	// Serialize the parameters back into a query string.
	var serializedParams []string
	for key, value := range params {
		// Format each key-value pair as "key=value" and append it to the list.
		serializedParams = append(serializedParams, fmt.Sprintf("%s=%v", key, value))
	}

	// Join all key-value pairs into a single query string.
	return strings.Join(serializedParams, "&"), nil
}
