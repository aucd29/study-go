package myrandstring

import "testing"

func TestRandomString(t *testing.T) {
	randString := NewRandString()
	randString.Init("burke")
	result := randString.Hello()

	switch result {
	case "Hi burke Welcome!",
		"Great to see you, burke!",
		"Hail, burke Well met!",
		"안녕 burke 환영해!":
		break

	default:
		t.Errorf("")
	}
}
