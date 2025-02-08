package errors

type HTTPError interface {
	error
	Code() ErrCode
	Message() string
	Status() int
}

type AppError struct {
	code    ErrCode
	message string
	err     error
	status  int
}

func NewAppError(code ErrCode, message string, err error) *AppError {
	return &AppError{
		code:    code,
		message: message,
		err:     err,
		status:  code.DefaultHTTPStatus(),
	}
}

func (e *AppError) Code() ErrCode {
	return e.code
}

func (e *AppError) Message() string {
	return e.message
}

func (e *AppError) Status() int {
	return e.status
}

func (e *AppError) Error() string {
	return e.message
}

func (e *AppError) Unwrap() error {
	return e.err
}
