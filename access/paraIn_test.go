package access

import (
	"testing"

	"github.com/llxxbb/platform-common/def"
	"github.com/stretchr/testify/assert"
)

func TestVerifierF_ParaNil(t *testing.T) {
	var p *ParaIn[string] = nil
	e := p.VerifyF(nil)
	assert.Equal(t, def.ET_BIZ, e.ErrType)
	assert.Equal(t, def.E_VERIFY.Code, e.Code)
	assert.Equal(t, def.E_VERIFY.Msg+" para can't be empty!", e.Msg)
	assert.Equal(t, nil, e.Context)
}
func TestVerifierF_DataNil(t *testing.T) {
	p := ParaIn[any]{}
	e := p.VerifyF(nil)
	assert.Equal(t, def.ET_BIZ, e.ErrType)
	assert.Equal(t, def.E_VERIFY.Code, e.Code)
	assert.Equal(t, def.E_VERIFY.Msg+" para can't be empty!", e.Msg)
	assert.Equal(t, nil, e.Context)
}
func TestVerifierF_FunNil(t *testing.T) {
	p := ParaIn[string]{Data: "lxb"}
	e := p.VerifyF(nil)
	assert.Nil(t, e)
}
func TestVerifierF_Fun(t *testing.T) {
	p := ParaIn[string]{Data: "lxb"}
	ce := def.CustomError{
		ErrorDefine: def.ErrorDefine{},
		ErrType:     def.ET_COM,
		Context:     nil,
	}
	e := p.VerifyF(func() *def.CustomError {
		return &ce
	})
	assert.Equal(t, ce, *e)
}
