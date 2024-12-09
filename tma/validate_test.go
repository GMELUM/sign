package tma

import "testing"

const (
	secret = `1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA`
	params = `user=%7B%22id%22%3A1093776793%2C%22first_name%22%3A%22%D0%90%D1%80%D1%82%D1%83%D1%80%22%2C%22last_name%22%3A%22%D0%A4%D1%80%D0%B0%D0%BD%D0%BA%22%2C%22username%22%3A%22gmelum%22%2C%22language_code%22%3A%22ru%22%2C%22is_premium%22%3Atrue%2C%22allows_write_to_pm%22%3Atrue%7D&chat_instance=3411281046910109270&chat_type=private&auth_date=1710181745&hash=ef19060b40a2277fa4debd9c6ad9b37b1e7ac1b6f467e53c66ca6d8df2c3c168`
)

func TestValidate(t *testing.T) {
	param, result := Validate(params, secret)
	if result == false {
		t.Errorf("signature verification failed")
	}

	t.Log(param)
}

func BenchmarkValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validate(params, secret)
	}
}
