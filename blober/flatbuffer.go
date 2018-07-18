package blober

import (
	"context"
	"io"
	"os"

	service "github.com/aybabtme/grpc-blob/gen/flatbuffer/service"
	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ service.BloberServer = (*FlatbufferGRPCBlober)(nil)

type FlatbufferGRPCBlober struct {
	FS Blober
}

func (b *FlatbufferGRPCBlober) Put(ctx context.Context, req *service.PutReq) (*flatbuffers.Builder, error) {
	fd, err := b.FS.Create(string(req.Name()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't create file: %v", err)
	}

	if _, err := fd.Write(req.BlobBytes()); err != nil {
		return nil, status.Errorf(codes.Internal, "can't write to file: %v", err)
	}
	if err := fd.Close(); err != nil {
		return nil, status.Errorf(codes.Internal, "can't close file: %v", err)
	}
	return flatbuffers.NewBuilder(0), nil
}

func (b *FlatbufferGRPCBlober) Stream(srv service.Blober_StreamServer) error {
	req, err := srv.Recv()
	if err == io.EOF {
		return nil
	}
	defer srv.SendAndClose(flatbuffers.NewBuilder(0))

	fd, err := b.FS.Create(string(req.Name()))
	if err != nil {
		return status.Errorf(codes.Internal, "can't create file: %v", err)
	}
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			if err := fd.Close(); err != nil {
				return status.Errorf(codes.Internal, "can't close file: %v", err)
			}
			return nil
		}
		if _, err := os.Stdout.Write(req.BlobBytes()); err != nil {
			return status.Errorf(codes.Internal, "can't write to file: %v", err)
		}
	}
}
