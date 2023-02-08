package def

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomError(t *testing.T) {
	ce := NewCustomError(ET_COM, 123, "hello", "lxb")
	assert.Equal(t, ET_COM, ce.ErrType)
	assert.Equal(t, 123, ce.Code)
	assert.Equal(t, "hello", ce.Msg)
	assert.Equal(t, "lxb", ce.Context)
}
