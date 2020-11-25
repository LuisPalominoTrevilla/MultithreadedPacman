package structures

// MessageBroker that can send and receive messages through channels
type MessageBroker struct {
	EatFood     chan struct{}
	PhaseChange chan int
}
