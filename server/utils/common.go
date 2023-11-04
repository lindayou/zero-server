package utils

type CommonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Body interface{} `json:"body"`
}

func SuccessResp(body interface{}) CommonResponse {
	return CommonResponse{200, "success", body}
}

func FailResp(body interface{}) CommonResponse {
	return CommonResponse{500, "error", body}
}
