package rabbit

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendMessage(t *testing.T) {
	assert := assert.New(t)
	messag, err := New("amqp://127.0.0.1/", map[string]struct{}{"test": {}})
	assert.NoError(err)
	data, err := json.Marshal(map[string]string{"tesqt": "asdasd"})
	assert.NoError(err)

	for i := 0; i < 10; i++ {
		err = messag.Put("test", data)
		assert.NoError(err)
	}

	messag.Close()
}
