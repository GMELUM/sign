package sign

import (
	"testing"
)

func TestCreate(t *testing.T) {

	var secret = "37VE8WK547U7A0Q49CRHKDNIHLHGOCTR"

	var item = make(map[string]interface{})
	item["item"] = "coin"
	item["count"] = 1000

	var params = make(map[string]interface{})
	params["user"] = 3
	params["order"] = 5
	params["task"] = 9
	params["items"] = []map[string]interface{}{item}

	token, err := Create(params, secret)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(token) <= 0 {
		t.Errorf("incorrent token")
	}

}

func BenchmarkCreate(b *testing.B) {

	var secret = "37VE8WK547U7A0Q49CRHKDNIHLHGOCTR"

	var params = make(map[string]interface{})
	params["type"] = "entry"
	params["count"] = 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Create(params, secret)
	}

}
