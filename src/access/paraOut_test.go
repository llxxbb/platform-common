package access

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.cdel.local/platform/go/platform-common/def"
)

func TestGetSuccessResult(t *testing.T) {
	rtn := GetSuccessResult[any]("lxb")
	assert.Equal(t, 0, rtn.State)
	assert.Equal(t, "lxb", rtn.Data)
}

func TestErrorResultED(t *testing.T) {
	rtn := GetErrorResultED[any](def.E_VERIFY, "lxb", "my test")
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
