package utils

import (
	"testing"
)

func TestIsJwtToken_ValidToken(t *testing.T) {
	validJwtToken := "valid.jwt.token"
	result := IsJwtToken(validJwtToken)

	if !result {
		t.Errorf("IsJwtToken() returned false for a valid JWT token")
	}
}

func TestIsJwtToken_InvalidToken(t *testing.T) {
	invalidToken := "invalid_token"
	result := IsJwtToken(invalidToken)

	if result {
		t.Errorf("IsJwtToken() returned true for an invalid token")
	}
}
