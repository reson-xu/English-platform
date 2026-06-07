package dto

type HealthResp struct {
	AppName string `json:"appName"`
	Status  string `json:"status"`
}

type DataResp struct {
	Data any `json:"data"`
}

type ErrorEnvelopeResp struct {
	Error ErrorResp `json:"error"`
}

type ErrorResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
