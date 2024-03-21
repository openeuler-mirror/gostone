// utils_test.go
package utils

import (
	"database/sql"
	"reflect"
	"testing"
	"work.ctyun.cn/git/GoStack/gostone/conf"
)

type TestLinkStruct struct {
	ID       string         `json:"id"`
	Enabled  int            `json:"enabled"`
	IsDomain int            `json:"is_domain"`
	ParentID sql.NullString `json:"parent_id"`
}

func TestSetSingleLink(t *testing.T) {
	testStruct := TestLinkStruct{
		ID:       "123",
		Enabled:  1,
		IsDomain: 1,
		ParentID: sql.NullString{String: "parent123", Valid: true},
	}

	ignoreFields := []string{"enabled"} // Fields to be ignored in the result map
	result := SetSingleLink(testStruct, UserPath, "user", ignoreFields)

	expectedResult := map[string]interface{}{
		"user": map[string]interface{}{
			"id":        "123",
			"is_domain": true,
			"parent_id": "parent123",
			"links": map[string]interface{}{
				"self": conf.Url + "/v3/users/123",
			},
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("SetSingleLink() returned incorrect result. Got: %+v, Expected: %+v", result, expectedResult)
	}
}
