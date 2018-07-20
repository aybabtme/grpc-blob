package blober_test

import (
	"bufio"
	"context"
	"io/ioutil"
	"net"
	"os"
	"testing"
	"time"

	"github.com/aybabtme/grpc-blob/blober"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func testBlober(t *testing.T, withClient func(func(blober.Blober))) {
	tests := []struct {
		name     string
		test     func(t testing.TB, client blober.Blober, name string, blob []byte)
		filename string
		blob     []byte
	}{
		{name: "Write", test: testBloberWrite, filename: "helloworld", blob: []byte("hellooooo world")},
		{name: "Read", test: testBloberRead, filename: "helloworld", blob: []byte("hellooooo world")},
		{name: "Put", test: testBloberPut, filename: "helloworld", blob: []byte("hellooooo world")},
		{name: "Get", test: testBloberGet, filename: "helloworld", blob: []byte("hellooooo world")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			withClient(func(client blober.Blober) {
				tt.test(t, client, tt.filename, tt.blob)
			})
		})
	}
}

func testBloberWrite(t testing.TB, client blober.Blober, name string, blob []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wc, err := client.Write(ctx, name)
	require.NoError(t, err)

	buf := bufio.NewWriter(wc)
	n, err := buf.Write(blob)
	require.NoError(t, err)
	require.Equal(t, len(blob), n)

	err = buf.Flush()

	require.NoError(t, err)

	require.NoError(t, wc.Close())

	got, err := client.Get(ctx, name)
	require.NoError(t, err)
	require.Equal(t, blob, got)
}

func testBloberRead(t testing.TB, client blober.Blober, name string, blob []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := client.Put(ctx, name, blob)
	require.NoError(t, err)

	rd, err := client.Read(ctx, name)
	require.NoError(t, err)
	defer func() {
		err = rd.Close()
		require.NoError(t, err)
	}()

	got, err := ioutil.ReadAll(rd)
	require.NoError(t, err)
	require.Equal(t, blob, got)
}

func testBloberPut(t testing.TB, client blober.Blober, name string, blob []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := client.Put(ctx, name, blob)
	require.NoError(t, err)

	err = client.Put(ctx, name, blob)
	require.NoError(t, err)
}

func testBloberGet(t testing.TB, client blober.Blober, name string, want []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	got, err := client.Get(ctx, name)
	require.True(t, os.IsNotExist(err))
	require.Nil(t, got)

	err = client.Put(ctx, name, want)
	require.NoError(t, err)

	got, err = client.Get(ctx, name)
	require.NoError(t, err)
	require.Equal(t, want, got)
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
