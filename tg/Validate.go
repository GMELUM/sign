package tg

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/gmelum/sign/types"
)

var hashSecret []byte

func Validate(params string, secret string) (types.TGUser, bool) {

	if hashSecret == nil {
		hash := sha256.New()
		hash.Write([]byte(secret))
		hashSecret = hash.Sum(nil)
	}

	query, err := url.ParseQuery(params)
	if err != nil {
		return types.TGUser{}, false
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
		return types.TGUser{}, false
	}

	sort.Strings(pairs)

	impHmac := hmac.New(sha256.New, hashSecret)
	impHmac.Write([]byte(strings.Join(pairs, "\n")))

	isValid := hex.EncodeToString(impHmac.Sum(nil)) == hash
	if isValid {

		IDstring := query.Get("id")
		IDint, err := strconv.Atoi(IDstring)
		if err != nil {
			IDint = 0
		}

		return types.TGUser{
			ID:        IDint,
			FirstName: query.Get("first_name"),
			LastName:  query.Get("last_name"),
			UserName:  query.Get("username"),
			PhotoURL:  query.Get("photo_url"),
		}, true
	}

	return types.TGUser{}, false

}
