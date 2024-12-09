package tg

import "testing"

func TestValidate(t *testing.T) {

	var (
		secret = `1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA`
		params = `auth_date=1693552072&first_name=Артур&id=1093776793&last_name=Франк&photo_url=https://t.me/i/userpic/320/F1-CaPdnOBWj3heb1ffplwJXTkie5mnU0QsJD1CGOQ4.jpg&username=gmelum&hash=3d5c5229fdf9781734c4c55353f2bc113d32f8fecb9d3ebadc4d4664de68f9a4`
	)

	param, result := Validate(params, secret)
	if result == false {
		t.Errorf("signature verification failed")
	}

	t.Log(param)
}

func BenchmarkValidate(b *testing.B) {

	var (
		secret = `1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA`
		params = `auth_date=1693552072&first_name=Артур&id=1093776793&last_name=Франк&photo_url=https://t.me/i/userpic/320/F1-CaPdnOBWj3heb1ffplwJXTkie5mnU0QsJD1CGOQ4.jpg&username=gmelum&hash=3d5c5229fdf9781734c4c55353f2bc113d32f8fecb9d3ebadc4d4664de68f9a4`
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Validate(params, secret)
	}
}
