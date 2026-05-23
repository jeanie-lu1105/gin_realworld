package security

import "testing"

func TestHashPassword(t *testing.T) {
	hashPassword, err := HashPassword("123456789")
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
		return
	}
	t.Logf("Hashed password: %s", hashPassword)

	check := ValidatePassword("123456789", hashPassword)
	if !check {
		t.Errorf("check password failed")
		return
	}
	t.Logf("Check password passed ")
}
