package ucerror

import "fmt"

type Error struct {
	Code    int
	Message string
	err     error
}

func NewError(m string, c int, err error) Error {
	return Error{
		Code:    c,
		Message: m,
		err:     err,
	}
}

func (e Error) Unwrap() error {
	return e.err
}

func (e Error) Error() string {
	if e.err != nil {
		return fmt.Sprintf("[ucerror_ERROR: %d] %s\n%v", e.Code, e.Message, e.err)
	} else {
		return fmt.Sprintf("[ucerror_ERROR: %d] %s", e.Code, e.Message)
	}
}
