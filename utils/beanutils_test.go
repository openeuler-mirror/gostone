package utils

import (
	"testing"
)

type SourceStruct struct {
	Name  string
	Value int
}

type TargetStruct struct {
	Name  string
	Value int
}

func TestCopyProperties(t *testing.T) {
	// Create instances of source and target structs
	source := SourceStruct{
		Name:  "John",
		Value: 42,
	}

	target := TargetStruct{}

	// Call CopyProperties with source and target structs
	CopyProperties(&target, source)

	// Assert that the values are copied correctly
	if target.Name != "John" || target.Value != 42 {
		t.Errorf("CopyProperties failed. Expected: {Name: \"John\", Value: 42}, Got: %+v", target)
	}
}
