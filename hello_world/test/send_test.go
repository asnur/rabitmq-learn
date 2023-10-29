package test

import (
	"hello_world"
	"testing"
)

func TestSend(t *testing.T) {
	hello_world.Send()
}

func TestReceive(t *testing.T) {
	hello_world.Receive()
}
