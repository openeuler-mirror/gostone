package policy

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestCheck(t *testing.T) {
	source := map[string]interface{}{
		"role":   []string{"admin"},
		"userId": "123",
	}
	target := map[string]interface{}{
		"userId": "12",
		"name":   "test",
	}
	assert.Equal(t, true, Check("get_user", source, target))
	assert.Equal(t, true, Check("get_project", source, target))
}
