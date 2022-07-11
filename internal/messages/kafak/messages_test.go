package kafak

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendMessage(t *testing.T) {
	assert := assert.New(t)
	messag, err := New("80.78.253.137:9092", "my-kafka-topic")
	assert.NoError(err)
	data, err := json.Marshal(map[string]string{"tesqt": "asdasd"})
	assert.NoError(err)

	for i := 0; i < 10; i++ {
		err = messag.Put("my-kafka-topic", data)
		assert.NoError(err)
	}
	for i := 0; i < 10; i++ {
		data_get, err := messag.Get("my-kafka-topic", "asd")
		assert.NoError(err)
		assert.Equal(data, data_get)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	messag.Close(ctx)
}
