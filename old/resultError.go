package old

import (
	"fmt"
	"strconv"

	"gitlab.cdel.local/platform/go/platform-common/def"
	"go.uber.org/zap"
)

var CustomErrorLogOut = false

type ResultError def.CustomError

func (e *ResultError) ToResult() *ServiceResult[any] {
	return ConvertError[any](e)
}

func ConvertError[T any](e *ResultError) *ServiceResult[T] {
	eMsg := fmt.Sprintf("ErrorType: %s, Code: %d, Msg: %s", e.ErrType, e.Code, e.Msg)
	if CustomErrorLogOut {
		switch e.ErrType {
		case def.ET_BIZ:
			zap.L().Warn(eMsg)
		case def.ET_ENV:
			zap.L().Warn(eMsg)
		case def.ET_COM:
			zap.L().Warn(eMsg)
		case def.ET_SYS:
			zap.L().Error(eMsg)
		}
	}
	result := ServiceResult[T]{}
	result.ErrorCode = strconv.Itoa(e.Code)
	result.ErrorMsg = e.Msg
	return &result
}

// SameError 进行类型转换但不改变错误信息。
func SameError[I any, O any](in *ServiceResult[I]) *ServiceResult[O] {
	result := ServiceResult[O]{}
	result.Success = in.Success
	result.ErrorMsg = in.ErrorMsg
	result.ErrorCode = in.ErrorCode
	return &result
}
