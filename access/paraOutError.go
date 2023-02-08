package access

import (
	"fmt"

	"gitlab.cdel.local/platform/go/platform-common/def"
	"go.uber.org/zap"
)

type ParaOutError def.CustomError

func (e *ParaOutError) ToParaOut() ParaOut {
	eMsg := fmt.Sprintf("ErrorType: %s, Code: %d, Msg: %s", e.ErrType, e.Code, e.Msg)
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
	result := ParaOut{}
	result.State = e.Code
	result.ErrType = e.ErrType
	result.ErrMsg = e.Msg
	result.Data = e.Context
	return result
}
