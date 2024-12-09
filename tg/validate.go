package tg

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"unsafe"

	"github.com/gmelum/sign/types"
	"github.com/gmelum/sign/utils"
)

// Validate checks parameters using the secret key and returns a TGUser object if everything is successful.
func Validate(params string, secret string) (types.TGUser, bool) {
	// Create the secret key (hash of bot_token).
	// This hash is used as the key for HMAC verification.
	hashSecret := sha256.Sum256([]byte(secret))

	// Initializing variables to store the data:
	// hash - stores the "hash" parameter passed in the URL for comparison.
	// pairs - array to hold other key-value pairs from the URL parameters (maximum size of 16).
	// pairsCount - keeps track of how many key-value pairs were added to the array.
	// buffer - a byte buffer to store the "data-check-string" (up to 512 bytes).
	// bufPos - tracks the current position in the buffer.
	var (
		hash       string
		pairs      [16]string // Assuming a maximum of 16 parameters.
		pairsCount int        // Counter for the number of parameters.
		buffer     [512]byte  // Buffer for the data-check-string.
		bufPos     int        // Current position in the buffer.
	)

	// Convert the parameter string to a byte slice for efficient processing.
	// We use unsafe.Pointer for fast type conversion from string to []byte.
	paramBytes := *(*[]byte)(unsafe.Pointer(&params)) // Fast conversion of string to byte slice.
	paramLen := len(paramBytes)

	// Parse the parameters manually by iterating through the byte slice.
	// This avoids using `Split` and gives full control over the parsing process.
	for i := 0; i < paramLen; {
		// Find the key by locating the '=' character.
		start := i
		for i < paramLen && paramBytes[i] != '=' {
			i++
		}
		if i == paramLen { // If '=' is not found, break the loop.
			break
		}
		key := string(paramBytes[start:i])
		i++ // Skip the '='.

		// Find the value by locating the '&' character (or end of string).
		start = i
		for i < paramLen && paramBytes[i] != '&' {
			i++
		}
		value := string(paramBytes[start:i])

		// If the key is "hash", store the hash value.
		// Otherwise, store the key-value pair for sorting later.
		if key == "hash" {
			hash = value
		} else {
			if pairsCount < len(pairs) {
				pairs[pairsCount] = key + "=" + value
				pairsCount++
			}
		}
		if i < paramLen {
			i++ // Skip the '&' character.
		}
	}

	// If the hash parameter is not found, the validation is unsuccessful.
	if hash == "" {
		return types.TGUser{}, false
	}

	// Sort the key-value pairs lexicographically by their keys to ensure correct order.
	utils.QuickSortStrings(pairs[:pairsCount])

	// Construct the "data-check-string" in the buffer.
	// This is the string that will be used for the HMAC computation.
	for i := 0; i < pairsCount; i++ {
		pair := pairs[i]
		copy(buffer[bufPos:], pair) // Copy the key-value pair into the buffer.
		bufPos += len(pair)
		if i < pairsCount-1 {
			buffer[bufPos] = '\n' // Add a newline separator between key-value pairs.
			bufPos++
		}
	}

	// Compute the HMAC-SHA256 of the "data-check-string".
	// HMAC ensures the integrity and authenticity of the data.
	impHmac := hmac.New(sha256.New, hashSecret[:])
	impHmac.Write(buffer[:bufPos])                       // Feed the buffer data into the HMAC function.
	computedHash := hex.EncodeToString(impHmac.Sum(nil)) // Convert the HMAC result to a hex string.

	// Compare the computed hash with the provided hash.
	// If they don't match, the request is invalid or has been tampered with.
	if computedHash != hash {
		return types.TGUser{}, false
	}

	// If the hashes match, parse the parameters and return a TGUser object.
	var user types.TGUser
	for i := 0; i < pairsCount; i++ {
		pair := pairs[i]
		eqIdx := -1
		// Find the index of the '=' character to split the key and value.
		for j := 0; j < len(pair); j++ {
			if pair[j] == '=' {
				eqIdx = j
				break
			}
		}
		if eqIdx == -1 {
			continue // Skip the pair if '=' is not found.
		}
		key := pair[:eqIdx]     // Extract the key.
		value := pair[eqIdx+1:] // Extract the value.

		// Set the corresponding field of the TGUser struct based on the key.
		switch key {
		case "id":
			user.ID, _ = strconv.Atoi(value) // Convert the user ID to an integer.
		case "first_name":
			user.FirstName = value // Set the user's first name.
		case "last_name":
			user.LastName = value // Set the user's last name.
		case "username":
			user.UserName = value // Set the user's username.
		case "photo_url":
			user.PhotoURL = value // Set the user's photo URL.
		}
	}

	// Return the TGUser object and true to indicate the validation was successful.
	return user, true
}

