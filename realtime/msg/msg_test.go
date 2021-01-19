package msg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProducer(t *testing.T) {
	data := &Msgbody{
		ID:     1,
		Sender: "Tom",
		Msg:    "Hello",
	}
	err := Producer(data)

	assert.Nil(t, err)
}
