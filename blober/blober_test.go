package blober_test

import (
	"bufio"
	"context"
	"testing"
	"time"

	"github.com/aybabtme/grpc-blob/blober"
	"github.com/stretchr/testify/require"
)

func testBloberCreate(t testing.TB, client blober.Blober, name string, payload []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	wc, err := client.Create(ctx, name)
	require.NoError(t, err)
	nn, err := bufio.NewWriter(wc).Write(payload)
	require.NoError(t, err)
	require.Equal(t, len(payload), nn)
	wc.Close(ctx)
}
