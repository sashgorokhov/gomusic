package structs

type ErrorRequestParam struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type ErrorStruct struct {
	Error_code int `json:"error_code"`
	Error_msg int `json:"error_msg"`
	Request_params []ErrorRequestParam `json:"request_params"`
}

type ErrorResponse struct {
	Error ErrorStruct `json:"error"`
}
