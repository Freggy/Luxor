package transport

type Producer interface {

	// Sends the given packet over the wire
	Produce(<-chan interface{})

	// Closes the connection and frees all resources.
	Close()
}
