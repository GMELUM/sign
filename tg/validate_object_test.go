package tg

import "testing"

func TestValidateObject(t *testing.T) {

	var (
		secret = `1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA`
		object = map[string]interface{}{
			"auth_date":  1693552072,
			"first_name": "Артур",
			"id":         1093776793,
			"last_name":  "Франк",
			"photo_url":  "https://t.me/i/userpic/320/F1-CaPdnOBWj3heb1ffplwJXTkie5mnU0QsJD1CGOQ4.jpg",
			"username":   "gmelum",
			"hash":       "3d5c5229fdf9781734c4c55353f2bc113d32f8fecb9d3ebadc4d4664de68f9a4",
		}
	)

	param, result := ValidateObject(object, secret)
	if result == false {
		t.Errorf("signature verification failed")
	}

	t.Log(param)
}

func BenchmarkValidateObject(b *testing.B) {

	var (
		secret = `1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA`
		object = map[string]interface{}{
			"auth_date":  1693552072,
			"first_name": "Артур",
			"id":         1093776793,
			"last_name":  "Франк",
			"photo_url":  "https://t.me/i/userpic/320/F1-CaPdnOBWj3heb1ffplwJXTkie5mnU0QsJD1CGOQ4.jpg",
			"username":   "gmelum",
			"hash":       "3d5c5229fdf9781734c4c55353f2bc113d32f8fecb9d3ebadc4d4664de68f9a4",
		}
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ValidateObject(object, secret)
	}
}
