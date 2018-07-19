package blober_test

import (
	"testing"

	"github.com/aybabtme/grpc-blob/blober"
	"github.com/aybabtme/grpc-blob/blober/blobersrv"
	service "github.com/aybabtme/grpc-blob/gen/flatbuffergrpc/service"
	flatbuffers "github.com/google/flatbuffers/go"

	"google.golang.org/grpc"
)

func TestFlatbufferGRPC(t *testing.T) {
	testBlober(t, func(fn func(blober.Blober)) {
		svc := &blobersrv.FlatbufferGRPCBlober{FS: blober.Memory()}
		cc, done := withGRPC(t,
			[]grpc.ServerOption{grpc.CustomCodec(flatbuffers.FlatbuffersCodec{})},
			[]grpc.DialOption{grpc.WithCodec(flatbuffers.FlatbuffersCodec{})},
			func(s *grpc.Server) { service.RegisterBloberServer(s, svc) },
		)
		defer done()

		client := blober.Flatbuffer(service.NewBloberClient(cc))

		fn(client)
	})
}
