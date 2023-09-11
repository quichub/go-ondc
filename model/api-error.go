package model

type ApiError struct {
	HttpStatus int
	Message    string
}

func (e ApiError) Error() string {
	return e.Message
}

func NewApiError(httpStatus int, message string) ApiError {
	return ApiError{httpStatus, message}
}
