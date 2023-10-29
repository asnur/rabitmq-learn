package test

import (
	"testing"
	"work_queues"
)

func TestSend(t *testing.T) {
	message := []string{"message.", "message..", "message...", "message....", "message....."}

	for _, msg := range message {
		work_queues.Send(msg)
	}
}
