package access

import (
	"testing"

	"github.com/llxxbb/platform-common/def"
	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	e := def.NewCustomError(def.ET_ENV, 123, "hello", nil)
	pe := ParaOutError(*e)
	rtn := pe.ToParaOut()
	assert.Equal(t, def.ET_ENV, rtn.ErrType)
	assert.Equal(t, 123, rtn.State)
	assert.Equal(t, "hello", rtn.ErrMsg)
	assert.Nil(t, rtn.Data)
}
