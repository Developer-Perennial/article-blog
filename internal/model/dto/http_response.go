package dto

type HttpResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func RespondSuccess(c int, data interface{}) *HttpResponse {
	return &HttpResponse{
		Status:  c,
		Message: "Success",
		Data:    data,
	}
}

func RespondError(c int, msg string) *HttpResponse {
	return &HttpResponse{
		Status:  c,
		Message: msg,
		Data:    nil,
	}
}
