package structures

// MessageBroker that can send and receive messages through channels
type MessageBroker struct {
	EatPellet          chan bool
	PowerPelletWoreOff chan struct{}
	PhaseChange        chan int
}
