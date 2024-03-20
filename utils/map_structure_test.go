// utils_test.go
package utils

import (
	"testing"
)

type TestMapStruct struct {
	Name  string
	Age   int
	Valid bool
}

func TestMap2Structure(t *testing.T) {
	testMap := map[string]interface{}{
		"Name":  "Gostone",
		"Age":   2,
		"Valid": true,
	}

	var testStruct TestMapStruct
	err := Map2Structure(testMap, &testStruct)

	if err != nil {
		t.Errorf("Map2Structure() returned an error: %v", err)
	}

	expectedStruct := TestMapStruct{
		Name:  "Gostone",
		Age:   2,
		Valid: true,
	}

	if testStruct != expectedStruct {
		t.Errorf("Map2Structure() did not map the values correctly. Got: %+v, Expected: %+v", testStruct, expectedStruct)
	}
}

func TestMap2Structure_InvalidInput(t *testing.T) {
	invalidMap := "invalid input"
	var testStruct TestMapStruct
	err := Map2Structure(invalidMap, &testStruct)

	if err == nil {
		t.Errorf("Map2Structure() did not return an error for invalid input")
	}
}
