package spf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullResolver_LookupTXT(t *testing.T) {
	var r NullResolver

	resp, err := r.LookupTXT("example.com")
	assert.NoError(t, err)
	assert.Equal(t, []string{"example.com"}, resp)
}

func TestRealResolver_LookupTXT(t *testing.T) {
	var r RealResolver

	resp, err := r.LookupTXT("example.net")
	assert.NoError(t, err)
	assert.Equal(t, []string{"v=spf1 -all"}, resp)
}

func TestErrorResolver_LookupTXT(t *testing.T) {
	var r ErrorResolver

	resp, err := r.LookupTXT("example.com")
	assert.Error(t, err)
	assert.Empty(t, resp)
}
