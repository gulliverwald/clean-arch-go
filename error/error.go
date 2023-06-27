package customError

type CustomError struct {
	error
	ErrorCode    int
	ErrorMessage string
}

type ErrorRepository interface {
	New(errCode int, errMessage string) *CustomError
	Error() string
}
