package dbError

type AlreadyRegisteredError struct {
	Message string
}

func (err AlreadyRegisteredError) Error() string {
	return err.Message
}
