package schema

type ValidationError struct {
	Message string `json:"message"`
}

func NewValidationError(message string) ValidationError {
	return ValidationError{message}
}

func (e ValidationError) Error() string {
	return e.Message
}

