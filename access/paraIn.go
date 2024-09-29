package access

import (
	"gitlab.cdel.local/platform/go/platform-common/def"
)

type ParaIn[T any] struct {
	Data     T                 `json:"data"`               // 要传送的业务数据
	TraceId  string            `json:"traceId,omitempty"`  // 用于链路分析
	PTraceId string            `json:"pTraceId,omitempty"` // 父 traceId，用于解决跟踪分叉问题
	RefId    string            `json:"refId,omitempty"`    // 关联业务的ID,用于审计
	Auth     *Auth             `json:"auth"`               // 授权相关信息
	Label    map[string]string `json:"label,omitempty"`    // 用于标记请求者的额外信息
	Note     string            `json:"note,omitempty"`     // 备注信息， 会存储于行为采集
	Time     string            `json:"time,omitempty"`     // 发起端的调用时间, Long 型的字符串
}

type Auth struct {
	SysId   int    `json:"sysId,omitempty"`   // 授权验证用于那个系统
	RoleId  int    `json:"roleId,omitempty"`  // 要操作的角色
	Domain  string `json:"domain,omitempty"`  // 要验证的领域
	Operate string `json:"operate,omitempty"` // 要对领域进行的操作
}

func CreateFrom[T any](data T) ParaIn[T] {
	rtn := ParaIn[T]{}
	rtn.Data = data
	return rtn
}

func (para *ParaIn[T]) Verify() *def.CustomError {
	return para.VerifyF(nil)
}

// VerifyF
// Used to verify imputed parameter
// If verified ok then return nil, otherwise return {@link ParaOut}
func (para *ParaIn[T]) VerifyF(verifier func() *def.CustomError) *def.CustomError {
	var data any
	if para == nil {
		data = nil
	} else {
		data = para.Data
	}
	if data == nil {
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
