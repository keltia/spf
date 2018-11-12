package spf

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDomain_Empty(t *testing.T) {
	d, err := NewDomain("")
	assert.Error(t, err)
	assert.Empty(t, d)
	assert.Nil(t, d.ctx)
}

func TestNewDomain_Good(t *testing.T) {
	d, err := NewDomain("example.net")
	d.ctx.r = NullResolver{}
	assert.NoError(t, err)
	assert.NotEmpty(t, d)
	assert.NotEmpty(t, d)
	assert.NotNil(t, d.ctx)
	assert.Equal(t, "example.net", d.Name)
}

func TestDomain_Fetch(t *testing.T) {
	d, err := NewDomain("example.net")
	require.NoError(t, err)
	require.NotNil(t, d)

	err = d.Fetch()
	assert.NoError(t, err)
	assert.NotEmpty(t, d.Raw)
}

func TestDomain_FetchError(t *testing.T) {
	d, err := NewDomain("example.net")
	require.NoError(t, err)
	require.NotNil(t, d)

	d.ctx = &Context{&ErrorResolver{}}

	err = d.Fetch()
	assert.Error(t, err)
	assert.Empty(t, d.Raw)
}

func TestDomain_FetchMX_Null(t *testing.T) {
	d, err := NewDomain("example.net")
	require.NoError(t, err)
	require.NotNil(t, d)

	d.ctx = &Context{&NullResolver{}}

	mxs, err := d.FetchMX()
	assert.NoError(t, err)
	assert.Empty(t, mxs)
}

func TestDomain_FetchMX_Error(t *testing.T) {
	d, err := NewDomain("example.net")
	require.NoError(t, err)
	require.NotNil(t, d)

	d.ctx = &Context{&ErrorResolver{}}

	mxs, err := d.FetchMX()
	assert.Error(t, err)
	assert.Empty(t, mxs)
}

func TestDomain_Unroll(t *testing.T) {
	d, err := NewDomain("example.net")
	require.NoError(t, err)
	require.NotNil(t, d)

	err = d.Fetch()
	assert.NoError(t, err)
	assert.NotEmpty(t, d.Raw)

	r, err := d.Unroll(0)
	assert.NoError(t, err)
	assert.Empty(t, r.IPs)
}

func TestDomain_Unroll_Empty(t *testing.T) {
	d := &Domain{Name:"example.net"}

	d.ctx = &Context{NullResolver{}}
	r, err := d.Unroll(0)
	assert.Empty(t, r)
	assert.Error(t, err)
}

func TestSetVerbose(t *testing.T) {
	assert.False(t, fVerbose)
	SetVerbose()
	assert.True(t, fVerbose)
	fVerbose = false
}

func TestSetDebug(t *testing.T) {
	assert.False(t, fDebug)
	SetDebug()
	assert.True(t, fDebug)
	assert.True(t, fVerbose)
	fDebug = false
	fVerbose = false
}

func TestReset(t *testing.T) {
	fVerbose = true
	Reset()
	require.False(t, fVerbose)
}

func TestVersion(t *testing.T) {
	require.Equal(t, myVersion, Version())
}
