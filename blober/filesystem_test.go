package blober_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/aybabtme/grpc-blob/blober"
	"github.com/stretchr/testify/require"
)

func TestFileSystem(t *testing.T) {
	testBlober(t, func(fn func(blober.Blober)) {
		name, err := ioutil.TempDir(os.TempDir(), "test")
		require.NoError(t, err)
		defer os.RemoveAll(name)

		client := blober.FileSystem(name)
		fn(client)
	})
}

func BenchmarkFileSystem(b *testing.B) {
	benchBlober(b, func(fn func(blober.Blober)) {
		name, err := ioutil.TempDir(os.TempDir(), "test")
		require.NoError(b, err)
		defer os.RemoveAll(name)

		client := blober.FileSystem(name)
		fn(client)
	})
}
