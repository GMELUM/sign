package sign

import (
	"testing"
)

func TestValidate(t *testing.T) {

	var secret = "37VE8WK547U7A0Q49CRHKDNIHLHGOCTR"
	var hash = "b574755cfbb1a253c92477367b5d77af69df8d59553d72b9f315a99ea975d5c4"

	var params = make(map[string]string)
	params["type"] = "entry"
	params["count"] = "1"

	isValid, err := Validate(params, secret, hash)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !isValid {
		t.Errorf("incorrect")
	}

}

func BenchmarkValidate(b *testing.B) {

	var secret = "37VE8WK547U7A0Q49CRHKDNIHLHGOCTR"
	var hash = "b574755cfbb1a253c92477367b5d77af69df8d59553d72b9f315a99ea975d5c4"

	var params = make(map[string]string)
	params["type"] = "entry"
	params["count"] = "1"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Validate(params, secret, hash)
	}

}
