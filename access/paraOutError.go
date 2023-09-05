package access

import (
	"fmt"

	"gitlab.cdel.local/platform/go/platform-common/def"
	"go.uber.org/zap"
)

var CustomErrorLogOut = false

type ParaOutError def.CustomError

func (e *ParaOutError) ToParaOut() *ParaOut[any] {
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
	result := ParaOut[any]{}
	result.State = e.Code
	result.ErrType = e.ErrType
	result.ErrMsg = e.Msg
	result.Data = e.Context
	return &result
}

func ConvertError[T any](e *ParaOutError) *ParaOut[T] {
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
	result := ParaOut[T]{}
	result.State = e.Code
	result.ErrType = e.ErrType
	result.ErrMsg = e.Msg
	return &result
}

// SameError 进行类型转换但不改变错误信息。
func SameError[I any, O any](in *ParaOut[I]) *ParaOut[O] {
	result := ParaOut[O]{}
	result.State = in.State
	result.ErrType = in.ErrType
	result.ErrMsg = in.ErrMsg
	result.WarnMsg = in.WarnMsg
	return &result
}
