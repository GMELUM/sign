package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/vmihailenco/msgpack/v5"
)

type Data struct {
	Data msgpack.RawMessage `msgpack:"data"`
	Hash string             `msgpack:"hash"`
}

func Create(data interface{}, secret string) ([]byte, error) {

	serializedData, err := msgpack.Marshal(data)
	if err != nil {
		return nil, err
	}

	h := hmac.New(sha256.New, []byte(secret))
	h.Write(serializedData)
	hash := hex.EncodeToString(h.Sum(nil))

	return msgpack.Marshal(Data{
		Data: serializedData,
		Hash: hash,
	})

}
