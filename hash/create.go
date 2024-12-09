package hash

import (
	"io"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"

	"github.com/vmihailenco/msgpack/v5"
)

func Create(data map[string]string, secret string) (string, error) {

	params, err := msgpack.Marshal(data)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(params))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], params)

	token := hex.EncodeToString(ciphertext)

	return token, nil

}
