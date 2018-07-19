package blober_test

import (
	"testing"

	"github.com/aybabtme/grpc-blob/blober"
	"github.com/aybabtme/grpc-blob/blober/blobersrv"
	service "github.com/aybabtme/grpc-blob/gen/gofastgrpc"
	"google.golang.org/grpc"
)

func TestGoFastGRPC(t *testing.T) {
	testBlober(t, func(fn func(blober.Blober)) {
		svc := &blobersrv.GoFastGRPCBlober{FS: blober.Memory()}
		cc, done := withGRPC(t, nil, nil, func(s *grpc.Server) { service.RegisterBloberServer(s, svc) })
		defer done()

		client := blober.GoFastGRPC(service.NewBloberClient(cc))

		fn(client)
	})
}
