package tma

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"sort"
	"strings"
)

var hashSecret []byte

type User struct {
	ID                    int    `json:"id"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	UserName              string `json:"username"`
	PhotoURL              string `json:"photo_url"`
	Language              string `json:"language_code"`
	IsPremium             bool   `json:"is_premium"`
	AllowsWriteToPM       bool   `json:"allows_write_to_pm"`
	AddedToAttachmentMenu bool   `json:"added_to_attachment_menu"`
}

type Params struct {
	User         *User      `json:"user"`
	ChatType     *string    `json:"chat_type"`
	ChatInstance *string    `json:"chat_instance"`
}

func Validate(params, secret string) (Params, bool) {

	if hashSecret == nil {
		hash := hmac.New(sha256.New, []byte("WebAppData"))
		hash.Write([]byte(secret))
		hashSecret = hash.Sum(nil)
	}

	query, err := url.ParseQuery(params)
	if err != nil {
		return Params{}, false
	}

	var (
		hash     string
		pairs    = make([]string, 0, len(query))
	)

	for k, v := range query {
		if k == "hash" {
			hash = v[0]
			continue
		}
		pairs = append(pairs, k+"="+v[0])
	}

	if hash == "" {
		return Params{}, false
	}

	sort.Strings(pairs)

	impHmac := hmac.New(sha256.New, hashSecret)
	impHmac.Write([]byte(strings.Join(pairs, "\n")))

	println(hex.EncodeToString(impHmac.Sum(nil)))

	isValid := hex.EncodeToString(impHmac.Sum(nil)) == hash
	if isValid {

		param := new(Params)

		userData := query.Get("user")
		if userData != "" {
			user := new(User)
			err := json.Unmarshal([]byte(userData), user)
			if err == nil {
				param.User = user
			}
		}

		chatType := query.Get("chat_type")
		if chatType != "" {
			param.ChatType = &chatType
		}

		chatInstance := query.Get("chat_instance")
		if chatInstance != "" {
			param.ChatInstance = &chatInstance
		}

		return *param, true

	}

	return Params{}, false

}
