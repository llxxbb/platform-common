package old

import (
	"context"

	"github.com/llxxbb/platform-common/def"
)

type ServiceResult[T any] struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
	Retry     bool   `json:"retry"`
	Result    T      `json:"result"`
}

func GetSuccess[T any](v T) ServiceResult[T] {
	result := ServiceResult[T]{}
	result.Result = v
	result.Success = true
	return result
}

func GetFailure[T any](errorCode string, msg string) ServiceResult[T] {
	result := ServiceResult[T]{}
	result.ErrorCode = errorCode
	result.ErrorMsg = msg
	return result
}

// GetResultWithParam
//
//	execute the fun with param and automatically return the result whether an Exception in there.
//	all so with log.
func GetResultWithParam[P any, T any](p P, fn func(p P) (T, *def.CustomError)) *ServiceResult[T] {
	if fn == nil {
		msg := def.E_UNKNOWN.Msg + "The param [fn] doesn't provide"
		return GetErrorResultD[T](def.ET_SYS, def.E_UNKNOWN.Code, msg, nil)
	}
	t, e := fn(p)
	if e != nil {
		return GetErrorResult[T](e)
	}
	return GetSuccessResult(t)
}

// GetResultByParaCtx
//
//	execute the fun by param and Context and automatically return the result whether an Exception in there.
//	all so with log.
func GetResultByParaCtx[P any, T any](c context.Context, p P, fn func(c context.Context, p P) (T, *def.CustomError)) *ServiceResult[T] {
	if fn == nil {
		msg := def.E_UNKNOWN.Msg + "The param [fn] doesn't provide"
		return GetErrorResultD[T](def.ET_SYS, def.E_UNKNOWN.Code, msg, nil)
	}
	t, e := fn(c, p)
	if e != nil {
		return GetErrorResult[T](e)
	}
	return GetSuccessResult(t)
}

func GetErrorResultD[T any](et def.ErrorType, code int, msg string, context any) *ServiceResult[T] {
	e := def.NewCustomError(et, code, msg, context)
	return GetErrorResult[T](e)
}

func GetSuccessResult[T any](v T) *ServiceResult[T] {
	result := ServiceResult[T]{}
	result.Success = true
	result.Result = v
	return &result
}

func GetErrorResult[T any](e *def.CustomError) *ServiceResult[T] {
	myE := ResultError(*e)
	return ConvertError[T](&myE)
}
