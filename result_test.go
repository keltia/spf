package spf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResult_AppendMX(t *testing.T) {

}

func TestResult_ParseTXT_Empty(t *testing.T) {
	r := Result{}
	f, err := r.parseTXT("")
	assert.Error(t, err)
	assert.Nil(t, f)
}

func TestResult_ParseTXT_Good(t *testing.T) {
	SetDebug()
	r := Result{}
	f, err := r.parseTXT("salesforce.com")
	assert.NoError(t, err)
	assert.NotEmpty(t, f)
	t.Logf("f=%s", f)
	t.Logf("rec=%d", r.rec)
	t.Logf("dns=%d", r.dns)
	assert.Equal(t, f, r.IPs)
	Reset()
}

func TestResult_ParseTXT_Keltia(t *testing.T) {
	SetDebug()
	r := Result{}
	f, err := r.parseTXT("keltia.net")
	assert.NoError(t, err)
	assert.NotEmpty(t, f)
	t.Logf("f=%s", f)
	t.Logf("rec=%d", r.rec)
	t.Logf("dns=%d", r.dns)
	assert.Equal(t, f, r.IPs)
	Reset()
}
