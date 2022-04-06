package use_case

import "fmt"

type UseCaseError struct {
	Code    int
	Message string
	err     error
}

func NewError(m string, c int, err error) *UseCaseError {
	return &UseCaseError{
		Code:    c,
		Message: m,
		err:     err,
	}
}

func (e *UseCaseError) Unwrap() error {
	return e.err
}

func (e *UseCaseError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("[USE_CASE_ERROR: %d] %s\n%v", e.Code, e.Message, e.err)
	} else {
		return fmt.Sprintf("[USE_CASE_ERROR: %d] %s", e.Code, e.Message)
	}
}
