package exception

type NotFoundError struct {
	Message  string
	Payload  interface{}
	CanRetry bool
}

func (e NotFoundError) Error() string {
	return e.Message
}

func (e NotFoundError) Data() any {
	return e.Payload
}

func (err NotFoundError) MustRetry() bool {
	return err.CanRetry
}
