package test

import (
	"os"
	"testing"
	"work_queues"
)

func TestReceive(t *testing.T) {
	args := []string{"message.", "message..", "message...", "message....", "message....."}

	for _, arg := range args {
		os.Args = append(os.Args, arg)
		work_queues.Worker()
	}
}
