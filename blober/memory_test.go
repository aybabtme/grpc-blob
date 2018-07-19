package blober_test

import (
	"testing"

	"github.com/aybabtme/grpc-blob/blober"
)

func TestMemory(t *testing.T) {
	testBlober(t, func(fn func(blober.Blober)) {
		client := blober.Memory()
		fn(client)
	})
}
