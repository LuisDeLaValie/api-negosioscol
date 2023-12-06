package utils

import "fmt"

type ErrorStatusCode struct {
	Message string `json:"error_description"`
	Code    int    `json:"errno"`
	Errorr  string `json:"error"`
}

func Error400(format string, a ...any) *ErrorStatusCode {
	return &ErrorStatusCode{
		Message: fmt.Sprintf(format, a...),
		Code:    400,
		Errorr:  "bad_request",
	}
}
func Error401(format string, a ...any) *ErrorStatusCode {
	return &ErrorStatusCode{
		Message: fmt.Sprintf(format, a...),
		Code:    401,
		Errorr:  "unauthorized",
	}
}
func Error403(format string, a ...any) *ErrorStatusCode {
	return &ErrorStatusCode{
		Message: fmt.Sprintf(format, a...),
		Code:    403,
		Errorr:  "forbidden",
	}
}
func Error404(format string, a ...any) *ErrorStatusCode {
	return &ErrorStatusCode{
		Message: fmt.Sprintf(format, a...),
		Code:    404,
		Errorr:  "not_found",
	}
}
func Error500(format string, a ...any) *ErrorStatusCode {
	return &ErrorStatusCode{
		Message: fmt.Sprintf(format, a...),
		Code:    500,
		Errorr:  "internal_error",
	}
}

func (e *ErrorStatusCode) Error() string {
	return fmt.Sprintf("Error %d %s:  %s", e.Code, e.Errorr, e.Message)
}
