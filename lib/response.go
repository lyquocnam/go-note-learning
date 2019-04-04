package lib

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *Response) Error() string {
	return r.Message
}

func NewResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewErrorReponse(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
	}
}
