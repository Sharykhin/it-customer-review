package contract

type (
	// MessageProvider
	MessageProvider interface {
		Listen() (<-chan []byte, error)
	}
)
