package utils

import (
	"encoding/hex"
	"testing"
)

func TestGenerateRandomBytes(t *testing.T) {
	size := 16
	randomBytes, err := GenerateRandomBytes(size)

	if err != nil {
		t.Errorf("GenerateRandomBytes() returned an error: %v", err)
	}

	if len(randomBytes) != size {
		t.Errorf("GenerateRandomBytes() returned bytes of incorrect size")
	}
}

func TestGenerateUUID(t *testing.T) {
	uuid := GenerateUUID()

	// Assuming a valid UUID is 32 characters long
	if len(uuid) != 32 {
		t.Errorf("GenerateUUID() returned an invalid UUID: %s", uuid)
	}
}

func TestFormatUUIDFromString(t *testing.T) {
	formattedUUID := FormatUUIDFromString("550e8400-e29b-41d4-a716-446655440000")
	expected := "550e8400e29b41d4a716446655440000"

	if formattedUUID != expected {
		t.Errorf("FormatUUIDFromString() returned %s, expected %s", formattedUUID, expected)
	}
}

func TestFormatUUID(t *testing.T) {
	buf, _ := hex.DecodeString("550e8400e29b41d4a716446655440000")
	formattedUUID, err := FormatUUID(buf)

	if err != nil {
		t.Errorf("FormatUUID() returned an error: %v", err)
	}

	expected := "550e8400e29b41d4a716446655440000"
	if formattedUUID != expected {
		t.Errorf("FormatUUID() returned %s, expected %s", formattedUUID, expected)
	}
}

func TestParseUUID(t *testing.T) {
	validUUID := "550e8400-e29b-41d4-a716-446655440000"
	parsedUUID, err := ParseUUID(validUUID)

	if err != nil {
		t.Errorf("ParseUUID() returned an error for a valid UUID: %v", err)
	}

	expectedBytes, _ := hex.DecodeString("550e8400e29b41d4a716446655440000")
	for i, b := range parsedUUID {
		if b != expectedBytes[i] {
			t.Errorf("Parsed UUID does not match expected bytes")
			break
		}
	}
}

func TestParseUUIDInvalid(t *testing.T) {
	invalidUUID := "invalid-uuid"
	_, err := ParseUUID(invalidUUID)

	if err == nil {
		t.Errorf("ParseUUID() did not return an error for an invalid UUID")
	}
}
