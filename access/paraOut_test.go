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
func TestGetResult_NoFun(t *testing.T) {
	rtn := GetResult[any](nil, "")
	assert.Equal(t, def.E_UNKNOWN.Code, rtn.State)
	assert.Equal(t, def.ErrorType("SYS"), rtn.ErrType)
	assert.Equal(t, def.E_UNKNOWN.Msg+"The param [fn] doesn't provide", rtn.ErrMsg)
	assert.Equal(t, nil, rtn.Data)
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
}

func TestGetSuccessResult(t *testing.T) {
	rtn := GetSuccessResult("lxb")
	assert.Equal(t, 0, rtn.State)
	assert.Equal(t, "lxb", rtn.Data)
}

func TestErrorResult(t *testing.T) {
	ce := def.CustomError{ErrorDefine: def.E_VERIFY, ErrType: "lxb", Context: "my test"}
	rtn := GetErrorResult[any](ce)
	assert.Equal(t, def.E_VERIFY.Code, rtn.State)
	assert.Equal(t, def.E_VERIFY.Msg, rtn.ErrMsg)
	assert.Equal(t, def.ErrorType("lxb"), rtn.ErrType)
}
