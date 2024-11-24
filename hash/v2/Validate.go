package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/vmihailenco/msgpack/v5"
)

func Validate(entity interface{}, data []byte, secret string) error {
	// Десериализуем данные в структуру Data
	var deserializeData Data
	if err := msgpack.Unmarshal(data, &deserializeData); err != nil {
		return err
	}

	// Генерируем хэш с использованием HMAC
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(deserializeData.Data)
	expectedHash := hex.EncodeToString(h.Sum(nil))

	// Сравниваем хэш с переданным
	if !hmac.Equal([]byte(deserializeData.Hash), []byte(expectedHash)) {
		return errors.New("invalid data")
	}

	// Десериализуем данные в конечную структуру
	if err := msgpack.Unmarshal(deserializeData.Data, entity); err != nil {
		return err
	}

	return nil
}