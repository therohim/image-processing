package exception

type ClientError struct {
	Message  string
	Payload  interface{}
	CanRetry bool
}

func (err ClientError) Error() string {
	return err.Message
}

func (err ClientError) Data() any {
	return err.Payload
}

func (err ClientError) MustRetry() bool {
	return err.CanRetry
}

type TooManyRequestError struct {
	Message  string
	Payload  interface{}
	CanRetry bool
}

func (err TooManyRequestError) Error() string {
	return err.Message
}

func (err TooManyRequestError) Data() any {
	return err.Payload
}

func (err TooManyRequestError) MustRetry() bool {
	return err.CanRetry
}
