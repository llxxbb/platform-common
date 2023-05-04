package http

import (
	"fmt"
	"gitlab.cdel.local/platform/go/platform-common/access"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	t.Skip("this is an example rather than a test")

	client := ClientNoPP(500, "http://server.for.test", &http.Client{})

	rtn := access.ParaOut[any]{}
	errRtn := Post(client, &access.ParaIn[any]{}, &rtn, "/subUrl")
	fmt.Print(errRtn.ErrMsg)
}
