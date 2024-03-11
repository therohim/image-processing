package exception

import "Ubersnap-middle-backend-programmer-test/utils"

type ErrorData interface {
	Error() string
	Data() any
	MustRetry() bool
}

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}

// Becareful using this function,
// because it will call panic immediately
func PanicAndLog(err ErrorData) {
	utils.Logger.Warn(err.Error(), utils.LogAny("error", err.Error()), utils.LogAny("data", err.Data()))
	panic(err)
}
