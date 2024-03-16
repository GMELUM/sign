package tma

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"sort"
	"strings"

	"github.com/gmelum/sign/types"
)

var hashSecret []byte

func Validate(params, secret string) (types.TMAUser, bool) {

	if hashSecret == nil {
		hash := hmac.New(sha256.New, []byte("WebAppData"))
		hash.Write([]byte(secret))
		hashSecret = hash.Sum(nil)
	}

	query, err := url.ParseQuery(params)
	if err != nil {
		return types.TMAUser{}, false
	}

	var (
		hash  string
		pairs = make([]string, 0, len(query))
	)

	for k, v := range query {
		if k == "hash" {
			hash = v[0]
			continue
		}
		pairs = append(pairs, k+"="+v[0])
	}

	if hash == "" {
		return types.TMAUser{}, false
	}

	sort.Strings(pairs)

	impHmac := hmac.New(sha256.New, hashSecret)
	impHmac.Write([]byte(strings.Join(pairs, "\n")))

	isValid := hex.EncodeToString(impHmac.Sum(nil)) == hash
	if isValid {

		user := types.TMAUser{}

		userData := query.Get("user")
		if userData != "" {
			err := json.Unmarshal([]byte(userData), &user)
			if err != nil {
				return user, false
			}
		}

		chatType := query.Get("chat_type")
		if chatType != "" {
			user.ChatType = chatType
		}

		chatInstance := query.Get("chat_instance")
		if chatInstance != "" {
			user.ChatInstance = chatInstance
		}

		return user, true

	}

	return types.TMAUser{}, false

}
