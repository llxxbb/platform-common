package access

import (
	"fmt"

	"gitlab.cdel.local/platform/go/platform-common/def"
	"go.uber.org/zap"
)

type ParaOut[T any] struct {
	State   int           `json:"state"`   // 等于0正常，小于0异常。大于0 警告
	Data    T             `json:"data"`    // 返回的业务数据。
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
 *
 * @param sur    : Function to be executed
 * @param errMsg : Only used for log
 * @param <T>    ：What Type would you like to return
 * @return : take the result into {@link ParaOut}
 */
func GetResult[T any](bd func() (T, *def.CustomError), errMsg string) ParaOut[T] {
	t, e := bd()
	if e != nil {
		e.Msg += errMsg
		eMsg := fmt.Sprintf("ErrorType: %s, Code: %d, Msg: %s", e.ErrType, e.Code, e.Msg)
		switch e.ErrType {
		case def.ET_BIZ:
			zap.L().Warn(eMsg)
		case def.ET_ENV:
			zap.L().Warn(eMsg)
		case def.ET_COM:
			zap.L().Warn(eMsg)
		default:
			zap.L().Error(fmt.Sprintf("ErrorType: %s, Code: %d, Msg: %s", e.ErrType, e.Code, e.Msg))
		}
		return GetErrorResult(e.Code, e.ErrType, e.Msg, e.Context.(T))
	}
	return GetSuccessResult(t)
}

func GetSuccessResult[T any](v T) ParaOut[T] {
	result := ParaOut[T]{}
	result.State = 0
	result.Data = v
	return result
}

func GetErrorResultE[T any](e def.CustomError) ParaOut[T] {
	result := ParaOut[T]{}
	result.State = e.Code
	result.ErrType = e.ErrType
	result.ErrMsg = e.Msg
	result.Data = e.Context.(T)
	return result
}

func GetErrorResult[T any](state int, errType def.ErrorType, errMsg string, d T) ParaOut[T] {
	result := ParaOut[T]{}
	result.State = state
	result.ErrType = errType
	result.ErrMsg = errMsg
	result.Data = d
	return result
}
