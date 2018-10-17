package spf

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchTXT_Null(t *testing.T) {
	ctx := &Context{NullResolver{}}
	txt, err := fetchTXT(ctx, "")
	assert.Error(t, err)
	assert.Empty(t, txt)
}

func TestFetchTXT_NilContext(t *testing.T) {
	txt, err := fetchTXT(nil, "example.net")
	assert.Error(t, err)
	assert.Empty(t, txt)
}

func TestFetchTXT_Error(t *testing.T) {
	ctx := &Context{ErrorResolver{}}
	txt, err := fetchTXT(ctx, "example.net")
	assert.Error(t, err)
	assert.Empty(t, txt)
}

func TestGetSPF_Null(t *testing.T) {
	require.Empty(t, getSPF(nil))
}

func TestGetSPF_None(t *testing.T) {
	require.Empty(t, getSPF([]string{"fake txt record"}))
}

func TestGetSPF_One(t *testing.T) {
	r := "v=spf1 -all"
	res := getSPF([]string{r})
	assert.NotEmpty(t, res)
	assert.Equal(t, res, r)
}

func TestGetSPF_Two(t *testing.T) {
	r1 := "v=spf1 second"
	r2 := "v=spf1 first"
	res := getSPF([]string{r1, r2})
	assert.NotEmpty(t, res)
	assert.Equal(t, r1, res)
}

func TestGetSPF_NotFirst(t *testing.T) {
	r1 := "foobar"
	r2 := "v=spf1 second"
	r3 := "v=spf1 first"
	res := getSPF([]string{r1, r2, r3})
	assert.NotEmpty(t, res)
	assert.Equal(t, r2, res)
}
