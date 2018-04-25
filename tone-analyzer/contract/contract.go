package contract

type (
	// MessageProvider describes general interface for listening messages
	MessageProvider interface {
		Listen() (<-chan []byte, error)
	}

	// Logger describes general methods for logging
	Logger interface {
		Errorf(format string, args ...interface{})
	}
)
