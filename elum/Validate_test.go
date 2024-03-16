package elum

import (
	"testing"

	"github.com/gmelum/sign/types"
)

func TestValidate(t *testing.T) {

	var Secret = "37VE8WK547U7A0Q49CRHKDNIHLHGOCTR"

	var params = types.EncodeParams{
		ID:       123,
		IP:       "127.0.0.1",
		Platform: types.VKID,
		Expires:  123,
		User: types.User{
			NID:       123,
			FirstName: "Artur",
			LastName:  "Frank",
			UserName:  "GMELUM",
			Language:  "EN",
			IsPremium: true,
		},
	}

	token, err := Create(&params, Secret)
	if err != nil {
		t.Errorf(err.Error())
	}

	data, isValid := Validate(token, Secret)
	if !isValid {
		t.Errorf("is not valid token")
	}

	if data.IP != "127.0.0.1" {
		t.Errorf("token data is not valid")
	}

	user, err := data.ParseUser()
	if err != nil {
		t.Errorf(err.Error())
	}

	if user.NID != 123 {
		t.Errorf("incorrent data user")
	}

}

func BenchmarkValidate(b *testing.B) {

	var Secret = "37VE8WK547U7A0Q49CRHKDNIHLHGOCTR"

	var params = types.EncodeParams{
		ID:       123,
		IP:       "127.0.0.1",
		Platform: types.VKID,
		Expires:  123,
		User: types.User{
			NID:       123,
			FirstName: "Artur",
			LastName:  "Frank",
			UserName:  "GMELUM",
			Language:  "EN",
			IsPremium: true,
		},
	}

	token, err := Create(&params, Secret)
	if err != nil {
		b.Errorf(err.Error())
	}

	for i := 0; i < b.N; i++ {
		Validate(token, Secret)
	}

}
