package hash

import (
	"testing"
)

func TestCreate(t *testing.T) {

	var Secret = "37VE8WK547U7A0Q49CRHKDNIHLHGOCTR"

	var params = make(map[string]string)
	params["name"] = "clan_1.jpg"

	token, err := Create(params, Secret)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(token) <= 0 {
		t.Errorf("incorrent token")
	}

}

func BenchmarkCreate(b *testing.B) {

	var Secret = "37VE8WK547U7A0Q49CRHKDNIHLHGOCTR"

	var params = make(map[string]string)
	params["name"] = "clan_1.jpg"

	for i := 0; i < b.N; i++ {
		Create(params, Secret)
	}

}
