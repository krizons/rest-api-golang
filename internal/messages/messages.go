package messages

type Messages interface {
	Close()
	Put(queue string, data []byte) error
}
