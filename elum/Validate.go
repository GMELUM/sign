package elum

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/vmihailenco/msgpack/v5"
)

func Validate(token, secret string, params interface{}) bool {

	data, err := hex.DecodeString(token)
	if err != nil {
		return false
	}

	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return false
	}

	if len(data) < aes.BlockSize {
		return false
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, data)

	err = msgpack.Unmarshal(data, params)

	return err == nil

}
