package dbError

type CustomError struct {
	Message string
}

func (ce CustomError) Error() string {
	return ce.Message
}
