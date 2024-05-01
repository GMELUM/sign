package sign

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"sort"
)

func Validate(data map[string]string, secret string, hash string) (bool, error) {

	var buffer bytes.Buffer

	keys := make([]string, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		buffer.WriteString(key)
		buffer.WriteString(data[key])
	}

	if hashSecret == nil {
		hash := sha256.New()
		hash.Write([]byte(secret))
		hashSecret = hash.Sum(nil)
	}

	impHmac := hmac.New(sha256.New, hashSecret)
	impHmac.Write(buffer.Bytes())

	return hex.EncodeToString(impHmac.Sum(nil)) == hash, nil

}
