package messages

import "context"

type Messages interface {
	Close(ctx context.Context) error
	Put(queue string, data []byte) error
}
