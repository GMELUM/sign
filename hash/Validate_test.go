package hash

import (
	"testing"
)

func TestValidate(t *testing.T) {

	var Secret = "37VE8WK547U7A0Q49CRHKDNIHLHGOCTR"

	var params = make(map[string]string)
	params["name"] = "clan_1.jpg"

	token, err := Create(params, Secret)
	if err != nil {
		t.Errorf(err.Error())
	}

	data, isValid := Validate(token, Secret)
	if !isValid {
		t.Errorf("is not valid token")
	}

	if data["name"] != "clan_1.jpg" {
		t.Errorf("token data is not valid")
	}

}

func BenchmarkValidate(b *testing.B) {

	var Secret = "37VE8WK547U7A0Q49CRHKDNIHLHGOCTR"

	var params = make(map[string]string)
	params["name"] = "clan_1.jpg"

	token, err := Create(params, Secret)
	if err != nil {
		b.Errorf(err.Error())
	}

	for i := 0; i < b.N; i++ {
		Validate(token, Secret)
	}

}
