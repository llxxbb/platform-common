package access

import (
	"gitlab.cdel.local/platform/go/platform-common/def"
)

type ParaOut struct {
	State   int           `json:"state"`   // 等于0正常，小于0异常。大于0 警告
	Data    any           `json:"data"`    // 返回的业务数据。
	ErrType def.ErrorType `json:"errType"` // 错误类类型，state < 0 时有意义。
	ErrMsg  string        `json:"errMsg"`  // 错误的具体信息，state < 0 时有意义。
	WarnMsg string        `json:"warnMsg"` // 警告信息。
}

type BizDataI interface {
	GetBizData() (any, *def.CustomError)
}

/**
 * Execute the fun and automatically the result whether an Exception in there.
 * all so with log.
 */
func GetResult(fn func() (any, *def.CustomError), errMsg string) ParaOut {
	if fn == nil {
		msg := def.E_UNKNOWN.Msg + "The param [fn] doesn't provide"
		return GetErrorResultD(def.ET_SYS, def.E_UNKNOWN.Code, msg, nil)
	}
	t, e := fn()
	if e != nil {
		return GetErrorResult(*e)
	}
	return GetSuccessResult(t)
}

func GetSuccessResult(v any) ParaOut {
	result := ParaOut{}
	result.State = 0
	result.Data = v
	return result
}

func GetErrorResult(e def.CustomError) ParaOut {
	myE := ParaOutError(e)
	return myE.ToParaOut()
}

func GetErrorResultD(et def.ErrorType, code int, msg string, context any) ParaOut {
	e := def.NewCustomError(et, code, msg, context)
	return GetErrorResult(e)
}
