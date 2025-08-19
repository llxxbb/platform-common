package http

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/llxxbb/platform-common/access"
	"github.com/llxxbb/platform-common/def"
	"github.com/llxxbb/platform-common/old"
	pphttp "github.com/pinpoint-apm/pinpoint-go-agent/plugin/http"
	"go.uber.org/zap"
)

func RpcClient(timeOut int, baseUrl string) *resty.Client {
	client := pphttp.WrapClient(nil) // pinpoint
	return ClientNoPP(timeOut, baseUrl, client)
}

// ClientNoPP Compared with upper: only no pinPoint. Can be used for testing
func ClientNoPP(timeOut int, baseUrl string, client *http.Client) *resty.Client {
	rtn := resty.NewWithClient(client)
	rtn.SetTimeout(time.Duration(timeOut) * time.Millisecond)
	rtn.SetHeader("Content-Type", "application/json")
	rtn.SetBaseURL(baseUrl)
	return rtn
}

func Post[I any, O any](client *resty.Client, input *access.ParaIn[I], rtn *access.ParaOut[O], url string) *access.ParaOut[O] {
	_, err := client.R().SetBody(input).SetResult(rtn).Post(url)
	if err != nil {
		zap.L().Warn(err.Error(), zap.String("url", url))
		msg := fmt.Sprintf("%s%s %s%s", client.BaseURL, url, def.ENV_M, err.Error())
		eRtn := access.GetErrorResultD[O](def.ET_ENV, def.ENV_C, msg, nil)
		return eRtn
	}
	return rtn
}

func WrappedPost[I any, O any](client *resty.Client, input I, url string) (O, *def.CustomError) {
	var rtn = access.ParaOut[O]{}
	in := access.ParaIn[I]{Data: input}
	_, err := client.R().SetBody(in).SetResult(&rtn).Post(url)
	if err != nil {
		msg := fmt.Sprintf("%s%s %s%s", client.BaseURL, url, def.ENV_M, err.Error())
		zap.L().Warn(msg)
		customError := def.NewCustomError(def.ET_ENV, def.ENV_C, msg, nil)
		return rtn.Data, customError
	}
	if rtn.State < 0 || rtn.State > 0 {
		msg := fmt.Sprintf("%s%s %s%s", client.BaseURL, url, def.ENV_M, rtn.ErrMsg)
		zap.L().Warn(msg)
		return rtn.Data, rtn.ToCustomError()
	}
	return rtn.Data, nil
}

func WrappedPostOld[I any, O any](client *resty.Client, input I, url string) (O, *def.CustomError) {
	var rtn = old.ServiceResult[O]{}
	in := old.Request[I]{Params: input}
	_, err := client.R().SetBody(in).SetResult(&rtn).Post(url)
	if err != nil {
		msg := fmt.Sprintf("%s%s %s%s", client.BaseURL, url, def.ENV_M, err.Error())
		zap.L().Warn(msg)
		customError := def.NewCustomError(def.ET_ENV, def.ENV_C, msg, nil)
		return rtn.Result, customError
	}
	if !rtn.Success {
		msg := fmt.Sprintf("%s%s %s %s%s", client.BaseURL, url, rtn.ErrorCode, def.ENV_M, rtn.ErrorMsg)
		zap.L().Warn(msg)
		atoi, _ := strconv.Atoi(rtn.ErrorCode)
		return rtn.Result, def.NewCustomError(def.ET_ENV, atoi, rtn.ErrorMsg, nil)
	}
	return rtn.Result, nil
}

func WrappedPostRaw[I any, O any](client *resty.Client, input I, url string) (O, *def.CustomError) {
	var rtn O
	_, err := client.R().SetBody(input).SetResult(&rtn).Post(url)
	if err != nil {
		msg := fmt.Sprintf("%s%s %s%s", client.BaseURL, url, def.ENV_M, err.Error())
		zap.L().Warn(msg)
		customError := def.NewCustomError(def.ET_ENV, def.ENV_C, msg, nil)
		return rtn, customError
	}
	return rtn, nil

}
