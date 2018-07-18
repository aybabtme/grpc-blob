package blobersrv

import (
	"context"
	"io"

	"github.com/aybabtme/grpc-blob/blober"
	service "github.com/aybabtme/grpc-blob/gen/flatbuffer/service"
	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ service.BloberServer = (*FlatbufferGRPCBlober)(nil)

type FlatbufferGRPCBlober struct {
	FS blober.Blober
}

func (b *FlatbufferGRPCBlober) Put(ctx context.Context, req *service.PutReq) (*flatbuffers.Builder, error) {
	fd, err := b.FS.Create(ctx, string(req.Name()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't create file: %v", err)
	}

	if _, err := fd.Write(ctx, req.BlobBytes()); err != nil {
		return nil, status.Errorf(codes.Internal, "can't write to file: %v", err)
	}
	if err := fd.Close(ctx); err != nil {
		return nil, status.Errorf(codes.Internal, "can't close file: %v", err)
	}
	return flatbuffers.NewBuilder(0), nil
}

func (b *FlatbufferGRPCBlober) Stream(srv service.Blober_StreamServer) error {
	ctx := srv.Context()
	req, err := srv.Recv()
	if err == io.EOF {
		return nil
	}
	defer srv.SendAndClose(flatbuffers.NewBuilder(0))

	fd, err := b.FS.Create(ctx, string(req.Name()))
	if err != nil {
		return status.Errorf(codes.Internal, "can't create file: %v", err)
	}
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			if err := fd.Close(ctx); err != nil {
				return status.Errorf(codes.Internal, "can't close file: %v", err)
			}
			return nil
		}
		if _, err := fd.Write(ctx, req.BlobBytes()); err != nil {
			return status.Errorf(codes.Internal, "can't write to file: %v", err)
		}
	}
}
