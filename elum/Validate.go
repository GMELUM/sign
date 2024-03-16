package elum

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/gmelum/sign/types"
	"github.com/vmihailenco/msgpack/v5"
)

func Validate(token, secret string) (types.DecodeParams, bool) {

	params := types.DecodeParams{}

	data, err := hex.DecodeString(token)
	if err != nil {
		return params, false
	}

	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return params, false
	}

	if len(data) < aes.BlockSize {
		return params, false
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, data)

	err = msgpack.Unmarshal(data, &params)
	if err != nil {
		return params, false
	}

	return params, true

}
