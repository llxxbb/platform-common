package old

type Request struct {
	Params       any    `json:"params"`
	Domain       string `json:"domain"`
	ResourcePath string `json:"resourcePath"`
	Platform     string `json:"platform"`
	AppType      string `json:"appType"`
	AppKey       string `json:"appKey"`
	AppVersion   string `json:"appVersion"`
	Sid          string `json:"sid"`
	UserID       string `json:"userID"`
	PublicKey    string `json:"publicKey"`
	AesKey       string `json:"aesKey"`
	CustomerIP   string `json:"customerIP"`
	Af           int8   `json:"af"`
	Fs           string `json:"fs"`
	Ve           string `json:"ve"`
	Lt           string `json:"lt"`
	Ap           string `json:"ap"`
	Time         string `json:"time"`
}
