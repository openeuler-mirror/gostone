package service

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestFormatUrl(t *testing.T) {
	url := "http://cinder-api.cty.os:10014/v2/%(project_id)s"
	assert.Equal(t, "http://cinder-api.cty.os:10014/v2/123", formatUrl(url, "123", "321"))
}
