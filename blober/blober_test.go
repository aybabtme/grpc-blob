package blober_test

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"testing"
	"time"

	"github.com/aybabtme/grpc-blob/blober"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func randbytes(size uint32) []byte {
	out := make([]byte, size)
	for i := range out {
		out[i] = byte(i % 256)
	}
	return out
}

func testBlober(t *testing.T, withClient func(func(blober.Blober))) {
	type test struct {
		name     string
		test     func(t testing.TB, client blober.Blober, name string, blob []byte, bufSize uint32)
		filename string
		blob     []byte
		bufSize  uint32
	}
	tests := []test{
		{name: "Write", test: testBloberWrite, filename: "helloworld", blob: []byte("hellooooo world"), bufSize: 1 << 16},
		{name: "Read", test: testBloberRead, filename: "helloworld", blob: []byte("hellooooo world"), bufSize: 1 << 16},
		{name: "Put", test: testBloberPut, filename: "helloworld", blob: []byte("hellooooo world"), bufSize: 1 << 16},
		{name: "Get", test: testBloberGet, filename: "helloworld", blob: []byte("hellooooo world"), bufSize: 1 << 16},
	}
	for byteSize := uint32(64); byteSize <= 1<<17; byteSize *= 2 {
		for bufSize := uint32(64); bufSize <= 1<<17; bufSize *= 2 {
			tests = append(tests, test{name: fmt.Sprintf("Read_large_%d_%d", byteSize, bufSize), test: testBloberRead, filename: "helloworld", blob: randbytes(byteSize), bufSize: bufSize})
			tests = append(tests, test{name: fmt.Sprintf("Write_large_%d_%d", byteSize, bufSize), test: testBloberWrite, filename: "helloworld", blob: randbytes(byteSize), bufSize: bufSize})
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			withClient(func(client blober.Blober) {
				tt.test(t, client, tt.filename, tt.blob, tt.bufSize)
			})
		})
	}
}

func testBloberWrite(t testing.TB, client blober.Blober, name string, blob []byte, bufSize uint32) {
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

func testBloberRead(t testing.TB, client blober.Blober, name string, blob []byte, bufSize uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := client.Put(ctx, name, blob)
	require.NoError(t, err)

	rd, err := client.Read(ctx, name, bufSize)
	require.NoError(t, err)
	defer func() {
		err = rd.Close()
		require.NoError(t, err)
	}()

	got, err := ioutil.ReadAll(rd)
	require.NoError(t, err)
	if !bytes.Equal(blob, got) {
		ioutil.WriteFile("got.raw", got, 0644)
		ioutil.WriteFile("want.raw", blob, 0644)
	}
	require.Equal(t, blob, got)
}

func testBloberPut(t testing.TB, client blober.Blober, name string, blob []byte, bufSize uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := client.Put(ctx, name, blob)
	require.NoError(t, err)

	err = client.Put(ctx, name, blob)
	require.NoError(t, err)
}

func testBloberGet(t testing.TB, client blober.Blober, name string, want []byte, bufSize uint32) {
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

func benchBlober(b *testing.B, withClient func(func(blober.Blober))) {
	type bench struct {
		name     string
		bench    func(b *testing.B, client blober.Blober, name string, blob []byte, bufSize uint32)
		filename string
		blob     []byte
		bufSize  uint32
	}
	var benchs []bench
	for byteSize := uint32(1 << 17); byteSize <= 1<<20; byteSize *= 2 {
		benchs = append(benchs,
			bench{name: fmt.Sprintf("Put_%d", byteSize), bench: benchBloberPut, filename: "helloworld", blob: randbytes(byteSize)},
			bench{name: fmt.Sprintf("Get_%d", byteSize), bench: benchBloberGet, filename: "helloworld", blob: randbytes(byteSize)},
			bench{name: fmt.Sprintf("Write_%d", byteSize), bench: benchBloberWrite, filename: "helloworld", blob: randbytes(byteSize)},
		)
		for bufSize := uint32(1 << 17); bufSize <= 1<<17; bufSize *= 2 {
			benchs = append(benchs, bench{name: fmt.Sprintf("Read_large_%d_%d", byteSize, bufSize), bench: benchBloberRead, filename: "helloworld", blob: randbytes(byteSize), bufSize: bufSize})
		}
	}
	for _, tt := range benchs {
		b.Run(tt.name, func(b *testing.B) {
			withClient(func(client blober.Blober) {
				tt.bench(b, client, tt.filename, tt.blob, tt.bufSize)
			})
		})
	}
}

func benchBloberPut(b *testing.B, client blober.Blober, name string, blob []byte, bufSize uint32) {
	b.SetBytes(int64(len(blob)))
	b.ReportAllocs()
	b.ResetTimer()
	err := client.Put(context.Background(), name, blob)
	if err != nil {
		b.Fatal(err)
	}
}
func benchBloberGet(b *testing.B, client blober.Blober, name string, blob []byte, bufSize uint32) {
	if err := client.Put(context.Background(), name, blob); err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(blob)))
	b.ReportAllocs()
	b.ResetTimer()
	got, err := client.Get(context.Background(), name)
	if err != nil {
		b.Fatal(err)
	}
	if len(blob) != len(got) {
		b.Fatalf("want=%d got=%d", len(blob), len(got))
	}
}
func benchBloberWrite(b *testing.B, client blober.Blober, name string, blob []byte, bufSize uint32) {
	fd, err := client.Write(context.Background(), name)
	if err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(blob)))
	b.ReportAllocs()
	b.ResetTimer()
	buf := bufio.NewWriter(fd)
	_, err = buf.Write(blob)
	if err != nil {
		b.Fatal(err)
	}
	err = buf.Flush()
	if err != nil {
		b.Fatal(err)
	}
	err = fd.Close()
	if err != nil {
		b.Fatal(err)
	}
}
func benchBloberRead(b *testing.B, client blober.Blober, name string, blob []byte, bufSize uint32) {
	if err := client.Put(context.Background(), name, blob); err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(blob)))
	b.ReportAllocs()
	b.ResetTimer()
	rd, err := client.Read(context.Background(), name, bufSize)
	if err != nil {
		b.Fatal(err)
	}
	_, err = io.Copy(ioutil.Discard, rd)
	if err != nil {
		b.Fatal(err)
	}
}

// helper

func withGRPC(t testing.TB, srvOpts []grpc.ServerOption, dialOpts []grpc.DialOption, configure func(s *grpc.Server)) (cc *grpc.ClientConn, done func()) {
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
