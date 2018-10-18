package spf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResult_AppendMX(t *testing.T) {

}

func TestResult_ParseTXT_Empty(t *testing.T) {
	r := Result{}
	err := r.parseTXT("")
	assert.Error(t, err)
	assert.Nil(t, r.IPs)
}

func TestResult_ParseTXT_Good(t *testing.T) {
	SetDebug()
	r := Result{}
	err := r.parseTXT("salesforce.com")
	assert.NoError(t, err)
	assert.NotEmpty(t, r.IPs)
	t.Logf("f=%s", r.IPs)
	t.Logf("rec=%d", r.rec)
	t.Logf("dns=%d", r.dns)
	Reset()
}

func TestResult_ParseTXT_Keltia(t *testing.T) {
	SetDebug()
	r := Result{}
	err := r.parseTXT("keltia.net")
	assert.NoError(t, err)
	assert.NotEmpty(t, r.IPs)
	t.Logf("f=%s", r.IPs)
	t.Logf("rec=%d", r.rec)
	t.Logf("dns=%d", r.dns)
	Reset()
}
