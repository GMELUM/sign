package sign

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

var hashSecret []byte

func MapToQueryString(data map[string]interface{}) string {
	var sb strings.Builder
	for k, v := range data {
	 sb.WriteString(fmt.Sprintf("%s=%v&", k, v))
	}
	return strings.TrimSuffix(sb.String(), "&")
   }

func Create(data map[string]interface{}, secret string) (string, error) {

	var buffer bytes.Buffer

	var keys []string
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		if key != "hash" {
			buffer.WriteString(key)
			buffer.WriteString(fmt.Sprintf("%v", data[key]))
		}
	}

	if hashSecret == nil {
		hash := sha256.New()
		hash.Write([]byte(secret))
		hashSecret = hash.Sum(nil)
	}

	impHmac := hmac.New(sha256.New, hashSecret)
	impHmac.Write(buffer.Bytes())

	return hex.EncodeToString(impHmac.Sum(nil)), nil

}
