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

func TestResult_ParseTXT(t *testing.T) {
	SetDebug()
	td := []string{"salesforce.com", "exacttarget.com", "keltia.net"}

	for _, d := range td {
		r := NewResult()
		err := r.parseTXT(d)
		assert.NoError(t, err)
		assert.NotEmpty(t, r.IPs)
		t.Logf("r=%s", r.String())
	}
	Reset()
}

func TestResult_ParseTXT_Recurs(t *testing.T) {
	r := &Result{rec: 100}
	err := r.parseTXT("example.net")

	assert.Error(t, err)
}
