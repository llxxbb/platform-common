package access

import (
	"gitlab.cdel.local/platform/go/platform-common/def"
)

type ParaIn struct {
	Data     any               `json:"data"`     // 要传送的业务数据
	TraceId  string            `json:"traceId"`  // 用于链路分析
	PTraceId string            `json:"pTraceId"` // 父 traceId，用于解决跟踪分叉问题
	RefId    string            `json:"refId"`    // 关联业务的ID,用于审计
	Label    map[string]string `json:"label"`    // 用于标记请求者的额外信息
	Note     string            `json:"note"`     // 备注信息， 会存储于行为采集
	Time     string            `json:"time"`     // 发起端的调用时间, Long 型的字符串
}

func CreateFrom(data any) ParaIn {
	rtn := ParaIn{}
	rtn.Data = data
	return rtn
}

func (para *ParaIn) Verify() *def.CustomError {
	return para.VerifyF(nil)
}

/**
 * Used to verify imputed parameter
 * If verified ok then return nil, otherwise return {@link ParaOut}
 */
func (para *ParaIn) VerifyF(verifier func() *def.CustomError) *def.CustomError {
	if para == nil || para.Data == nil {
		return &def.CustomError{
			ErrorDefine: def.ErrorDefine{
				Code: def.E_VERIFY.Code,
				Msg:  def.E_VERIFY.Msg + " para can't be empty!",
			},
			ErrType: def.ET_BIZ,
		}
	}
	if verifier == nil {
		return nil
	}
	return verifier()
}
