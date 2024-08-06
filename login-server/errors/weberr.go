package errors

// WebError represents a custom error with HTTP status code and error message
type WebError struct {
	StatusCode int    // HTTP status code
	Message    string // Error message
}

func (e *WebError) Error() string {
	return e.Message
}

func (e *WebError) Status() int {
	return e.StatusCode
}

func NewError(code int, message string) *WebError {
	err := &WebError{
		StatusCode: code,
		Message:    message,
	}
	return err
}
