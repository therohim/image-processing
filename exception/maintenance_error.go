package exception

type MaintenanceError struct {
	Message  string
	Payload  interface{}
	CanRetry bool
}

func (err MaintenanceError) Error() string {
	return err.Message
}

func (err MaintenanceError) Data() any {
	return err.Payload
}

func (err MaintenanceError) MustRetry() bool {
	return err.CanRetry
}
