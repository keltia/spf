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

func TestDomain_Unroll(t *testing.T) {
	d, err := NewDomain("example.net")
	require.NoError(t, err)
	require.NotNil(t, d)

	err = d.Fetch()
	assert.NoError(t, err)
	assert.NotEmpty(t, d.Raw)

	b, err := d.Unroll(0)
	assert.Nil(t, b)
	assert.NoError(t, err)
}

func TestVersion(t *testing.T) {
	require.Equal(t, myVersion, Version())
}
