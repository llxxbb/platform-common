package http

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	pphttp "github.com/pinpoint-apm/pinpoint-go-agent/plugin/http"
	"gitlab.cdel.local/platform/go/platform-common/access"
	"gitlab.cdel.local/platform/go/platform-common/def"
	"go.uber.org/zap"
	"net/http"
	"time"
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
		zap.L().Warn(err.Error(), zap.String("url", url))
		msg := fmt.Sprintf("%s%s %s%s", client.BaseURL, url, def.ENV_M, err.Error())
		customError := def.NewCustomError(def.ET_ENV, def.ENV_C, msg, nil)
		return rtn.Data, customError
	}
	if rtn.State < 0 {
		return rtn.Data, toErr(client.BaseURL, &rtn, url)
	}
	return rtn.Data, nil
}

func toErr[T any](baseUrl string, rtn *access.ParaOut[T], url string) *def.CustomError {
	if rtn.ErrType == def.ET_BIZ {
		// 自己的错误
		zap.L().Error(rtn.ErrMsg, zap.String("url", url))
		msg := fmt.Sprintf("%s%s %s%s", baseUrl, url, def.SYS_M, rtn.ErrMsg)
		return def.NewCustomError(def.ET_SYS, def.SYS_C, msg, nil)
	} else {
		// 别人的错误
		zap.L().Warn(rtn.ErrMsg, zap.String("url", url))
		msg := fmt.Sprintf("%s%s %s%s", baseUrl, url, def.ENV_M, rtn.ErrMsg)
		return def.NewCustomError(def.ET_ENV, def.ENV_C, msg, nil)
	}
}
