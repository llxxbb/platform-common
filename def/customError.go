package def

type ErrorDefine struct {
	Code int
	Msg  string
}
type CustomError struct {
	ErrorDefine
	ErrType ErrorType
	Context any
}

func NewCustomError(et ErrorType, code int, msg string, context any) CustomError {
	return CustomError{
		ErrorDefine: ErrorDefine{
			Code: code,
			Msg:  msg,
		},
		ErrType: et,
		Context: context,
	}
}

// call out err
var E_CALL_OPTION = ErrorDefine{-11, "call_option should not be null!"}
var E_CALL_PARA = ErrorDefine{-12, "call out para can't convert to json!"}
var E_CALL_SIGN = ErrorDefine{-13, "can't sign for the call: "}
var E_CALL_ENCRYPT = ErrorDefine{-14, "can't encrypt for the call: "}

// call returned err
var E_RESPONSE_SIGN = ErrorDefine{-21, "response error, can't verify the sign: "}
var E_RESPONSE_DECRYPT = ErrorDefine{-22, "response error, can't decrypt input: "}
var E_RESPONSE_BODY_EMPTY = ErrorDefine{-31, "response error, empty body!"}
var E_RESPONSE_ERROR = ErrorDefine{-32, "returned error: "}

// common error
var E_VERIFY = ErrorDefine{-41, "verify failed: "}
var E_ENV = ErrorDefine{-51, "env error: "}
var E_SYS = ErrorDefine{-61, "sys error: "}
var E_UNKNOWN = ErrorDefine{-99, "undefined error: "}
