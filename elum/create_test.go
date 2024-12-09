package elum

import (
	"testing"

	"github.com/gmelum/sign/types"
)

func TestCreate(t *testing.T) {

	var Secret = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"

	var params = types.EncodeParams{
		ID:       1,
		IP:       "127.0.0.1",
		Platform: types.TG,
		Expires:  123,
		User: types.User{
			NID:       1093776793,
			FirstName: "Artur",
			LastName:  "Frank",
			UserName:  "gmelum",
			Language:  "EN",
			IsPremium: true,
		},
	}

	token, err := Create(&params, Secret)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(token) <= 0 {
		t.Errorf("incorrent token")
	}
	
}

func BenchmarkCreate(b *testing.B) {

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

	for i := 0; i < b.N; i++ {
		Create(&params, Secret)
	}

}
