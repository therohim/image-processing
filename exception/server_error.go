package exception

type ServerError struct {
	Message  string
	Payload  interface{}
	CanRetry bool
}

func (err ServerError) Error() string {
	return err.Message
}

func (err ServerError) Data() any {
	return err.Payload
}

func (err ServerError) MustRetry() bool {
	return err.CanRetry
}
