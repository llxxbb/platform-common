package old

type ServiceResult[T any] struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
	Retry     bool   `json:"retry"`
	Result    T      `json:"result"`
}

func GetSuccess[T any](v T) ServiceResult[T] {
	result := ServiceResult[T]{}
	result.Result = v
	result.Success = true
	return result
}

func GetFailure[T any](errorCode string, msg string) ServiceResult[T] {
	result := ServiceResult[T]{}
	result.ErrorCode = errorCode
	result.ErrorMsg = msg
	return result
}
