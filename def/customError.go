package def

import "fmt"

type ErrorDefine struct {
	Code int
	Msg  string
}
type CustomError struct {
	ErrorDefine
	ErrType ErrorType
	Context any
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("type: %s, code: %d, msg: %s", e.ErrType, e.Code, e.Msg)
}

func NewCustomError(et ErrorType, code int, msg string, context any) *CustomError {
	return &CustomError{
		ErrorDefine: ErrorDefine{
			Code: code,
			Msg:  msg,
		},
		ErrType: et,
		Context: context,
	}
}

const (
	// call out err
	CALL_OPTION_C = -11
	CALL_OPTION_M = "call_option should not be null!"

	CALL_PARA_C = -12
	CALL_PARA_M = "call out para can't convert to json!"

	CALL_SIGN_C = -13
	CALL_SIGN_M = "can't sign for the call: "

	CALL_ENCRYPT_C = -14
	CALL_ENCRYPT_M = "can't encrypt for the call: "

	// call returned err
	RESPONSE_SIGN_C = -21
	RESPONSE_SIGN_M = "response error, can't verify the sign: "

	RESPONSE_DECRYPT_C = -22
	RESPONSE_DECRYPT_M = "response error, can't decrypt input: "

	RESPONSE_BODY_EMPTY_C = -31
	RESPONSE_BODY_EMPTY_M = "response error, empty body!"

	RESPONSE_ERROR_C = -32
	RESPONSE_ERROR_M = "returned error: "

	// common error
	VERIFY_C = -41
	VERIFY_M = "verify failed: "

	ENV_C = -51
	ENV_M = "env error: "

	SYS_C = -61
	SYS_M = "sys error: "

	UNKNOWN_C = -99
	UNKNOWN_M = "undefined error: "
)

// call out err
var E_CALL_OPTION = ErrorDefine{CALL_OPTION_C, CALL_OPTION_M}
var E_CALL_PARA = ErrorDefine{CALL_PARA_C, CALL_PARA_M}
var E_CALL_SIGN = ErrorDefine{CALL_SIGN_C, CALL_SIGN_M}
var E_CALL_ENCRYPT = ErrorDefine{CALL_ENCRYPT_C, CALL_ENCRYPT_M}

// call returned err
var E_RESPONSE_SIGN = ErrorDefine{RESPONSE_SIGN_C, RESPONSE_SIGN_M}
var E_RESPONSE_DECRYPT = ErrorDefine{RESPONSE_DECRYPT_C, RESPONSE_DECRYPT_M}
var E_RESPONSE_BODY_EMPTY = ErrorDefine{RESPONSE_BODY_EMPTY_C, RESPONSE_BODY_EMPTY_M}
var E_RESPONSE_ERROR = ErrorDefine{RESPONSE_ERROR_C, RESPONSE_ERROR_M}

// common error
var E_VERIFY = ErrorDefine{VERIFY_C, VERIFY_M}
var E_ENV = ErrorDefine{ENV_C, ENV_M}
var E_SYS = ErrorDefine{SYS_C, SYS_M}
var E_UNKNOWN = ErrorDefine{UNKNOWN_C, UNKNOWN_M}
