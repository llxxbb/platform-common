package http

import (
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"gitlab.cdel.local/platform/go/platform-common/access"
	"gitlab.cdel.local/platform/go/platform-common/def"
	"net/http"
	"testing"
)

func Test_Post_envErr(t *testing.T) {
	rtn := access.ParaOut[any]{}
	errRtn := Post(client, &access.ParaIn[any]{}, &rtn, "/subUrl")
	assert.Equal(t, def.ET_ENV, errRtn.ErrType)
	assert.Equal(t, "http://server.for.test/subUrl env error: Post \"http://server.for.test/subUrl\": dial tcp: lookup server.for.test: no such host", errRtn.ErrMsg)
}

func Test_Post_Ok(t *testing.T) {
	gock.New("").
		Post("/subUrl").
		Reply(http.StatusOK).
		JSON(map[string]any{
			"state": 0,
			"data":  "OK",
		})
	rtn := access.ParaOut[string]{}
	err := Post(client, &access.ParaIn[any]{}, &rtn, "/subUrl")
	assert.NotNil(t, err)
	assert.Equal(t, "OK", rtn.Data)
}

func Test_WrappedPost_realEnvErr(t *testing.T) {
	rtn, err := WrappedPost[string, string](client, "prj", "/subUrl")
	assert.Equal(t, "", rtn)
	assert.Equal(t, def.ET_ENV, err.ErrType)
	assert.Equal(t, "http://server.for.test/subUrl env error: Post \"http://server.for.test/subUrl\": gock: cannot match any request", err.Msg)
}

func Test_WrappedPost_returnEnvErr(t *testing.T) {
	gock.New("").
		Post("/subUrl").
		Reply(http.StatusOK).
		JSON(map[string]any{
			"state":   -1,
			"errType": "ENV",
			"errMsg":  "you get an err",
			"data":    "OK",
		})
	rtn, err := WrappedPost[string, string](client, "prj", "/subUrl")
	assert.Equal(t, "OK", rtn)
	assert.Equal(t, def.ET_ENV, err.ErrType)
	assert.Equal(t, "http://server.for.test/subUrl env error: you get an err", err.Msg)
}

func Test_WrappedPost_returnSysErr(t *testing.T) {
	gock.New("").
		Post("/subUrl").
		Reply(http.StatusOK).
		JSON(map[string]any{
			"state":   -1,
			"errType": "SYS",
			"errMsg":  "you get an err",
			"data":    "OK",
		})
	rtn, err := WrappedPost[string, string](client, "prj", "/subUrl")
	assert.Equal(t, "OK", rtn)
	assert.Equal(t, def.ET_ENV, err.ErrType)
	assert.Equal(t, "http://server.for.test/subUrl env error: you get an err", err.Msg)
}

func Test_WrappedPost_returnBizErr(t *testing.T) {
	gock.New("").
		Post("/subUrl").
		Reply(http.StatusOK).
		JSON(map[string]any{
			"state":   -1,
			"errType": "BIZ",
			"errMsg":  "you get an err",
			"data":    "OK",
		})
	rtn, err := WrappedPost[string, string](client, "prj", "/subUrl")
	assert.Equal(t, "OK", rtn)
	assert.Equal(t, def.ET_SYS, err.ErrType)
	assert.Equal(t, "http://server.for.test/subUrl sys error: you get an err", err.Msg)
}

func Test_WrappedPost_returnOk(t *testing.T) {
	gock.New("").
		Post("/subUrl").
		Reply(http.StatusOK).
		JSON(map[string]any{
			"state": 0,
			"data":  "OK",
		})
	rtn, err := WrappedPost[string, string](client, "prj", "/subUrl")
	assert.Equal(t, "OK", rtn)
	assert.Nil(t, err)
}

var client = ClientNoPP(500, "http://server.for.test", &http.Client{})

func TestMain(m *testing.M) {
	// before test
	gock.InterceptClient(client.GetClient())
	defer gock.Off()
	defer gock.RestoreClient(client.GetClient())
	// test
	m.Run()
	// after test
}
