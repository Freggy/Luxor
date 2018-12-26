package transport

// We could apply a nasty hack to only have one consumer for all packets.
// To do this we would have to create a Packet interface which only defines the method GetID(uint8)
// The next step is to define the id attribute in our protobuf files.
// Because the protobuf compiler will generate a GetID method for the field, all the conditions for
// the protobuf struct to implement the Packet interface would be satisfied.
type Consumer interface {

	// Consume starts listening for packets.
	Consume() chan<- Packet

	// Close closes the connection and frees all resources.
	Close()
}
