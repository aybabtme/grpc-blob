package blober_test

import (
	"testing"

	"github.com/aybabtme/grpc-blob/blober"
	"github.com/aybabtme/grpc-blob/blober/blobersrv"
	service "github.com/aybabtme/grpc-blob/gen/golanggrpc"
	"google.golang.org/grpc"
)

func TestGolangGRPC(t *testing.T) {
	testBlober(t, func(fn func(blober.Blober)) {
		svc := &blobersrv.GolangGRPCBlober{FS: blober.Memory()}
		cc, done := withGRPC(t, nil, nil, func(s *grpc.Server) { service.RegisterBloberServer(s, svc) })
		defer done()

		client := blober.GolangGRPC(service.NewBloberClient(cc))

		fn(client)
	})
}

func BenchmarkGolangGRPC(b *testing.B) {
	benchBlober(b, func(fn func(blober.Blober)) {
		svc := &blobersrv.GolangGRPCBlober{FS: blober.Memory()}
		cc, done := withGRPC(b, nil, nil, func(s *grpc.Server) { service.RegisterBloberServer(s, svc) })
		defer done()

		client := blober.GolangGRPC(service.NewBloberClient(cc))

		fn(client)
	})
}
