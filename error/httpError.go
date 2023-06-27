package customError

type HttpError struct {
	customError *CustomError
}

func NewHttpError() ErrorRepository {
	return &HttpError{
		customError: nil,
	}
}

func (err *HttpError) New(errCode int, errMessage string) *CustomError {
	err.customError = &CustomError{
		ErrorCode:    errCode,
		ErrorMessage: errMessage,
	}

	return err.customError
}

func (err *HttpError) Error() string {
	return err.customError.ErrorMessage
}
