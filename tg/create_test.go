package tg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {

	var (
		secret = "1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
		params = map[string]interface{}{
			"id":         123456789,
			"first_name": "John",
			"last_name":  "Doe",
			"username":   "johndoe",
			"photo_url":  "http://example.com/photo.jpg",
		}
	)

	// Call the Create function to generate the signed parameters.
	result, err := Create(params, secret)

	// Check for any errors during the function execution.
	assert.NoError(t, err)

	// Assert that the result is what we expect.
	// Ensure that the result string contains all parameters and the hash.
	assert.Contains(t, result, "first_name=John")
	assert.Contains(t, result, "id=123456789")
	assert.Contains(t, result, "last_name=Doe")
	assert.Contains(t, result, "photo_url=http://example.com/photo.jpg")
	assert.Contains(t, result, "username=johndoe")

	// Make sure the hash parameter is included, though we don't know the exact value here.
	// In an actual test, you would compare the hash with the expected computed value.
	assert.Contains(t, result, "hash=")

	_, valid := Validate(result, secret)

	assert.True(t, valid, true)

}
