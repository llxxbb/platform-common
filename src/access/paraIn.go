package access

type ParaIn[T any] struct {
	Data     T                 `json:"data"`     // 要传送的业务数据
	TraceId  string            `json:"traceId"`  // 用于链路分析
	PTraceId string            `json:"pTraceId"` // 父 traceId，用于解决跟踪分叉问题
	RefId    string            `json:"refId"`    // 关联业务的ID,用于审计
	Label    map[string]string `json:"label"`    // 用于标记请求者的额外信息
	Note     string            `json:"note"`     // 备注信息， 会存储于行为采集
	Time     string            `json:"time"`     // 发起端的调用时间, Long 型的字符串
}

func CreateFrom[T any](data *T) ParaIn[T] {
	rtn := ParaIn[T]{}
	rtn.Data = *data
	return rtn
}

// func Verify[T, U any](para *ParaIn[T]) ParaOut[U] {
// 	return VerifyS(para, null)
// }

/**
 * Used to verify imputed parameter
 * If verified ok then return null, otherwise return {@link ParaOut}
 *
 * @param para         : inputted data
 * @param dataVerifier : verify for data self, if null then return null
 */
// func VerifyS[T, U any] (para *ParaIn[T] , dataVerifier func() (ParaOut[U], error) ) ParaOut[U]{
// 	if para == nil || para.data == nil{
// 		return ParaOut.GetErrorResult(E_VERIFY.Code, E_BIZ, E_VERIFY.Msg + " para can't be empty!");
// 	}
// 	if dataVerifier == nil {
// 		return ParaOut[U]{}
// 	}
// 	try {
// 		return dataVerifier.get();
// 	} catch (Throwable e) {
// 		return ParaOut.getErrorResult(UNKNOWN_C, ErrorType.BIZ, UNKNOWN_M);
// 	}
// }
