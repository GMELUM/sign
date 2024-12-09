package tg

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strconv"
	"strings"

	"github.com/gmelum/sign/types"
)

// ValidateObject checks parameters using the secret key and returns a TGUser object if everything is successful.
func ValidateObject(params map[string]interface{}, secret string) (types.TGUser, bool) {
	// Create the secret key (hash of bot_token).
	hashSecret := sha256.Sum256([]byte(secret))

	// Variables for hash and key-value pairs.
	var hash string
	pairs := make([]string, 0, len(params)-1) // Allocate memory upfront for all pairs except "hash".

	// Process the map and prepare key-value pairs.
	for key, value := range params {
		if key == "hash" {
			hash, _ = value.(string) // Avoid unnecessary checks since "hash" is always a string.
		} else {
			// Convert the value to a string.
			valueStr := ""
			switch v := value.(type) {
			case string:
				valueStr = v
			case int:
				valueStr = strconv.Itoa(v)
			case float64:
				valueStr = strconv.FormatFloat(v, 'f', -1, 64)
			default:
				valueStr = ""
			}

			// Append the key-value pair to the slice.
			if valueStr != "" {
				pairs = append(pairs, key+"="+valueStr)
			}
		}
	}

	// If the hash parameter is missing, validation fails.
	if hash == "" {
		return types.TGUser{}, false
	}

	// Sort the pairs lexicographically.
	sort.Strings(pairs)

	// Construct the data-check string using a strings.Builder for efficiency.
	var dataCheckBuilder strings.Builder
	for i, pair := range pairs {
		if i > 0 {
			dataCheckBuilder.WriteByte('\n') // Add newline between pairs.
		}
		dataCheckBuilder.WriteString(pair)
	}
	dataCheckString := dataCheckBuilder.String()

	// Compute the HMAC-SHA256 of the data-check string.
	hmacFunc := hmac.New(sha256.New, hashSecret[:])
	hmacFunc.Write([]byte(dataCheckString))
	computedHash := hex.EncodeToString(hmacFunc.Sum(nil))

	// Compare computed hash with the provided hash.
	if computedHash != hash {
		return types.TGUser{}, false
	}

	// Parse the parameters into a TGUser object.
	var user types.TGUser
	for key, value := range params {
		switch key {
		case "id":
			if id, ok := value.(float64); ok {
				user.ID = int(id)
			}
		case "first_name":
			user.FirstName, _ = value.(string)
		case "last_name":
			user.LastName, _ = value.(string)
		case "username":
			user.UserName, _ = value.(string)
		case "photo_url":
			user.PhotoURL, _ = value.(string)
		}
	}

	return user, true
}
