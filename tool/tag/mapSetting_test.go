package tag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getMapSetting(t *testing.T) {
	// case: empty tag
	rtn := getMapSetting("")
	assert.Equal(t, "", rtn.key)
	assert.Equal(t, false, rtn.omit)
	assert.Equal(t, false, rtn.drillSub)

	// omit
	rtn = getMapSetting("hello,omitempty")
	assert.Equal(t, "hello", rtn.key)
	assert.Equal(t, true, rtn.omit)
	assert.Equal(t, false, rtn.drillSub)

	// drill sub
	rtn = getMapSetting("hello,sub")
	assert.Equal(t, "hello", rtn.key)
	assert.Equal(t, false, rtn.omit)
	assert.Equal(t, true, rtn.drillSub)
}
