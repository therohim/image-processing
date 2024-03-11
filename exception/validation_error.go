package exception

type ValidationError struct {
	Message  string
	Payload  interface{}
	CanRetry bool
}

func (err ValidationError) Error() string {
	return err.Message
}

func (err ValidationError) Data() any {
	return err.Payload
}

func (err ValidationError) MustRetry() bool {
	return err.CanRetry
}
