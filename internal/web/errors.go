package web

type ErrorCode string

const (
	ErrorNotFound    ErrorCode = "NOT_FOUND"
	ErrorInternal    ErrorCode = "INTERNAL"
	ErrorBadRequest  ErrorCode = "BAD_REQUEST"
	ErrorInvalidData ErrorCode = "INVALID_DATA"
	ErrorTimeout     ErrorCode = "TIMEOUT"
)

type ErrorBody struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}

func getErrResp(code ErrorCode, message string) *ErrorResponse {
	return &ErrorResponse{
		Error: ErrorBody{
			Code:    code,
			Message: message,
		},
	}
}
