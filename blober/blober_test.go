package blober_test

import (
	"bufio"
	"context"
	"net"
	"testing"
	"time"

	"github.com/aybabtme/grpc-blob/blober"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func testBlober(t *testing.T, withClient func(func(blober.Blober))) {
	tests := []struct {
		name     string
		test     func(t testing.TB, client blober.Blober, name string, payload []byte)
		filename string
		payload  []byte
	}{
		{name: "Write", test: testBloberWrite, filename: "helloworld", payload: []byte("hellooooo world")},
		{name: "Put", test: testBloberPut, filename: "helloworld", payload: []byte("hellooooo world")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			withClient(func(client blober.Blober) {
				tt.test(t, client, tt.filename, tt.payload)
			})
		})
	}
}

func testBloberWrite(t testing.TB, client blober.Blober, name string, payload []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wc, err := client.Write(ctx, name)
	require.NoError(t, err)

	buf := bufio.NewWriter(wc)
	nn, err := buf.Write(payload)
	require.NoError(t, err)
	require.Equal(t, len(payload), nn)

	err = buf.Flush()

	require.NoError(t, err)

	require.NoError(t, wc.Close())
}

func testBloberPut(t testing.TB, client blober.Blober, name string, payload []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := client.Put(ctx, name, payload)
	require.NoError(t, err)
}

// helper

func withGRPC(t *testing.T, srvOpts []grpc.ServerOption, dialOpts []grpc.DialOption, configure func(s *grpc.Server)) (cc *grpc.ClientConn, done func()) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	s := grpc.NewServer(srvOpts...)
	configure(s)
	go s.Serve(l)

	dialOpts = append([]grpc.DialOption{grpc.WithInsecure()}, dialOpts...)

	cc, err = grpc.Dial(l.Addr().String(), dialOpts...)
	require.NoError(t, err)

	return cc, func() {
		cc.Close()
		s.Stop()
		l.Close()
	}
}
