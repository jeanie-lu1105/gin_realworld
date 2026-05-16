package security

import "testing"

func TestHashPassword(t *testing.T) {
	hashPassword, err := HashPassword("secret123")
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
		return
	}
	t.Logf("Hashed password: %s", hashPassword)

	check := ValidatePassword("secret123", hashPassword)
	if !check {
		t.Errorf("check password failed")
		return
	}
	t.Logf("Check password passed", check)
}
