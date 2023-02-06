package access

import (
	"gitlab.cdel.local/platform/go/platform-common/def"
)

type ParaOut[T any] struct {
	State   int           `json:"state"`   // 等于0正常，小于0异常。大于0 警告
	Data    T             `json:"data"`    // 返回的业务数据。
	ErrType def.ErrorType `json:"errType"` // 错误类类型，state < 0 时有意义。
	ErrMsg  string        `json:"errMsg"`  // 错误的具体信息，state < 0 时有意义。
	WarnMsg string        `json:"warnMsg"` // 警告信息。
}

type BizData interface {
	GetBizData() (any error)
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
// func GetResult[T any](bd BizData, errMsg string) ParaOut[T]{
//     t, e := bd.GetBizData()
// 	if e != nil {
//         log.warn("eCode:{}, eType:{}, eMsg:{}, {}", e.code, e.eType, e.getMessage(), errMsg, e);
//         return GetErrorResult(e.code, e.eType, e.getMessage());
// 	}
//     try {
//     } catch (PlatformException e) {
//         log.warn("eCode:{}, eType:{}, eMsg:{}, {}", e.code, e.eType, e.getMessage(), errMsg, e);
//         return ParaOut.getErrorResult(e.code, e.eType, e.getMessage());
//     } catch (Throwable throwable) {
//         log.error("{}", errMsg, throwable);
//         return ParaOut.getErrorResult(UNKNOWN_C, ErrorType.SYS, UNKNOWN_M);
//     }
//     return getSuccessResult(t);
// }

func GetSuccessResult[T any](v T) ParaOut[T] {
	result := ParaOut[T]{}
	result.State = 0
	result.Data = v
	return result
}

func GetErrorResultED[T any](e def.ErrorDefine, et def.ErrorType, d T) ParaOut[T] {
	result := ParaOut[T]{}
	result.State = e.Code
	result.ErrType = et
	result.ErrMsg = e.Msg
	result.Data = d
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
