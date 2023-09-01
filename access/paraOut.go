package access

import (
	"gitlab.cdel.local/platform/go/platform-common/def"
)

type ParaOut[T any] struct {
	State   int           `json:"state"`             // 等于0正常，小于0异常。大于0 警告
	Data    T             `json:"data,omitempty"`    // 返回的业务数据。
	ErrType def.ErrorType `json:"errType,omitempty"` // 错误类类型，state < 0 时有意义。
	ErrMsg  string        `json:"errMsg,omitempty"`  // 错误的具体信息，state < 0 时有意义。
	WarnMsg string        `json:"warnMsg,omitempty"` // 警告信息。
}

type BizDataI interface {
	GetBizData() (any, *def.CustomError)
}

/**
 * Execute the fun and automatically the result whether an Exception in there.
 * all so with log.
 */
func GetResult[T any](fn func() (T, *def.CustomError), errMsg string) *ParaOut[T] {
	if fn == nil {
		msg := def.E_UNKNOWN.Msg + "The param [fn] doesn't provide"
		return GetErrorResultD[T](def.ET_SYS, def.E_UNKNOWN.Code, msg, nil)
	}
	t, e := fn()
	if e != nil {
		return GetErrorResult[T](*e)
	}
	return GetSuccessResult(t)
}

func GetSuccessResult[T any](v T) *ParaOut[T] {
	result := ParaOut[T]{}
	result.State = 0
	result.Data = v
	return &result
}

func GetErrorResult[T any](e def.CustomError) *ParaOut[T] {
	myE := ParaOutError(e)
	return ConvertError[T](&myE)
}

func GetErrorResultD[T any](et def.ErrorType, code int, msg string, context any) *ParaOut[T] {
	e := def.NewCustomError(et, code, msg, context)
	return GetErrorResult[T](e)
}
