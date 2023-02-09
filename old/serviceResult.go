package old

type ServiceResult struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
	Retry     bool   `json:"retry"`
	Result    any    `json:"result"`
}

func GetSuccess(v any) ServiceResult {
	result := ServiceResult{}
	result.Result = v
	result.Success = true
	return result
}

func GetFailure(errorCode string, msg string) ServiceResult {
	result := ServiceResult{}
	result.ErrorCode = errorCode
	result.ErrorMsg = msg
	return result
}
