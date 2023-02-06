package access

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.cdel.local/platform/go/platform-common/def"
)

func TestGetResult_Ok(t *testing.T) {
	f := func() (any, *def.CustomError) { return "lxb", nil }
	rtn := GetResult(f, "")
	assert.Equal(t, 0, rtn.State)
	assert.Equal(t, "lxb", rtn.Data)
}
func TestGetResult_Err(t *testing.T) {
	f := func() (any, *def.CustomError) {
		e := def.CustomError{
			ErrorDefine: def.ErrorDefine{Code: 5, Msg: "my err"}, ErrType: def.ET_BIZ, Context: "lxb",
		}
		return "", &e
	}
	rtn := GetResult(f, "")
	assert.Equal(t, 5, rtn.State)
	assert.Equal(t, def.ErrorType("BIZ"), rtn.ErrType)
	assert.Equal(t, "my err", rtn.ErrMsg)
	assert.Equal(t, "lxb", rtn.Data)
}

func TestGetSuccessResult(t *testing.T) {
	rtn := GetSuccessResult[any]("lxb")
	assert.Equal(t, 0, rtn.State)
	assert.Equal(t, "lxb", rtn.Data)
}

func TestErrorResultE(t *testing.T) {
	ce := def.CustomError{ErrorDefine: def.E_VERIFY, ErrType: "lxb", Context: "my test"}
	rtn := GetErrorResultE[any](ce)
	assert.Equal(t, def.E_VERIFY.Code, rtn.State)
	assert.Equal(t, def.E_VERIFY.Msg, rtn.ErrMsg)
	assert.Equal(t, def.ErrorType("lxb"), rtn.ErrType)
	assert.Equal(t, "my test", rtn.Data)
}

func TestErrorResult(t *testing.T) {
	rtn := GetErrorResult[any](5, "hello", "lxb", "my test")
	assert.Equal(t, 5, rtn.State)
	assert.Equal(t, def.ErrorType("hello"), rtn.ErrType)
	assert.Equal(t, "lxb", rtn.ErrMsg)
	assert.Equal(t, "my test", rtn.Data)
}
