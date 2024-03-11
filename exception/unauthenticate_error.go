package exception

type UnauthenticatedError struct {
	Message  string
	Payload  interface{}
	CanRetry bool
}

func (e UnauthenticatedError) Error() string {
	return e.Message
}

func (e UnauthenticatedError) Data() any {
	return e.Payload
}

func (err UnauthenticatedError) MustRetry() bool {
	return err.CanRetry
}
