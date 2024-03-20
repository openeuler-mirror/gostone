package utils

import (
	"encoding/json"
	"testing"
	"time"
)

func TestByte2Struct(t *testing.T) {
	type TestStruct struct {
		Name string
		Age  int
	}

	// Create a JSON byte slice
	jsonBytes := []byte(`{"Name":"John","Age":30}`)

	// Create an instance of TestStruct to unmarshal into
	var testObj TestStruct

	// Call Byte2Struct with the JSON byte slice
	Byte2Struct(jsonBytes, &testObj)

	// Assert that the values are as expected
	if testObj.Name != "John" || testObj.Age != 30 {
		t.Errorf("Byte2Struct failed. Expected: {\"Name\":\"John\",\"Age\":30}, Got: %+v", testObj)
	}
}

func TestStruct2Json(t *testing.T) {
	type TestStruct struct {
		Name string
		Age  int
	}

	// Create an instance of TestStruct
	testObj := TestStruct{
		Name: "Alice",
		Age:  25,
	}

	// Call Struct2Json with the TestStruct instance
	jsonString := Struct2Json(testObj)

	// Create expected JSON string
	expectedJSON := `{"Name":"Alice","Age":25}`

	// Assert that the JSON strings match
	if jsonString != expectedJSON {
		t.Errorf("Struct2Json failed. Expected: %s, Got: %s", expectedJSON, jsonString)
	}
}

func TestJSONRFC3339Milli_UnmarshalJSON(t *testing.T) {
	// Create a JSON string with RFC3339Milli formatted time
	jsonString := `"2022-02-19T12:34:56.789000Z"`

	// Create an instance of JSONRFC3339Milli to unmarshal into
	var jt JSONRFC3339Milli

	// Unmarshal the JSON string
	err := json.Unmarshal([]byte(jsonString), &jt)

	// Assert that there is no error and the time matches
	if err != nil {
		t.Errorf("JSONRFC3339Milli UnmarshalJSON failed. Error: %v", err)
	}

	expectedTime, _ := time.Parse(RFC3339Milli, "2022-02-19T12:34:56.789000Z")
	if time.Time(jt) != expectedTime {
		t.Errorf("JSONRFC3339Milli UnmarshalJSON failed. Expected: %v, Got: %v", expectedTime, time.Time(jt))
	}
}

func TestJSONRFC3339Milli_MarshalJSON(t *testing.T) {
	// Create a time instance with RFC3339Milli formatted time
	timeInstance := time.Date(2022, 2, 19, 12, 34, 56, 000000000, time.UTC)

	// Create an instance of JSONRFC3339Milli
	jt := JSONRFC3339Milli(timeInstance)

	// Marshal the JSONRFC3339Milli instance
	jsonBytes, err := json.Marshal(jt)

	// Assert that there is no error and the JSON string matches
	if err != nil {
		t.Errorf("JSONRFC3339Milli MarshalJSON failed. Error: %v", err)
	}

	expectedJSON := `"2022-02-19T12:34:56.000000Z"`
	if string(jsonBytes) != expectedJSON {
		t.Errorf("JSONRFC3339Milli MarshalJSON failed. Expected: %s, Got: %s", expectedJSON, string(jsonBytes))
	}
}

// Similar tests can be written for JSONRFC3339MilliNoZ
