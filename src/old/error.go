package old

type ErrorDefine struct {
	Code string
	Num  int
	Desc string
}

var E_PARAM_ERROR = ErrorDefine{"-1", -1, "请求参数错误"}
var E_UNKNOWN_RESOURCE = ErrorDefine{"-2", -2, "未知请求资源，无法找到对应接口"}
var E_ACCESS_TIMEOUT = ErrorDefine{"-3", -3, "请求超时"}
var E_ACCESS_ERROR = ErrorDefine{"-4", -4, "接口访问错误"}
var E_INTERNAL_ERROR = ErrorDefine{"-5", -5, "未知内部错误"}
var E_PUBLICKEY_TIMEOUT = ErrorDefine{"-6", -6, "请求公钥已过期"}
var E_NEED_LOGIN = ErrorDefine{"-7", -7, "需要登录"}
var E_NEED_AUTHORIZED = ErrorDefine{"-8", -8, "需要被授权"}
var E_FMT_ILLEGAL = ErrorDefine{"-9", -9, "请求数据格式非法"}
var E_RESOURCE_UNDEFINED = ErrorDefine{"-10", -10, "资源未定义"}
var E_OVER_THREAD_NUM = ErrorDefine{"-11", -11, "超过请求数，请稍后再试"}
var E_UNKNOWN_AES_KEY = ErrorDefine{"-12", -12, "不符合要求的密钥"}
var E_REQUEST_TIMEOUT = ErrorDefine{"-13", -13, "请求已过期"}
var E_DECODE_ERROR = ErrorDefine{"-14", -14, "解密失败"}
var E_DUPLICATED_REQUEST = ErrorDefine{"-15", -15, "重复请求"}
var E_BREAKER = ErrorDefine{"-16", -16, "进入熔断"}
