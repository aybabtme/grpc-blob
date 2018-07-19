package blobersrv

import (
	"context"
	"io"

	"github.com/aybabtme/grpc-blob/blober"
	service "github.com/aybabtme/grpc-blob/gen/flatbuffergrpc/service"
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

	if _, err := fd.Write(req.BlobBytes()); err != nil {
		return nil, status.Errorf(codes.Internal, "can't write to file: %v", err)
	}
	if err := fd.Close(); err != nil {
		return nil, status.Errorf(codes.Internal, "can't close file: %v", err)
	}
	fb := flatbuffers.NewBuilder(0)
	service.PutResStart(fb)
	fb.Finish(service.PutResEnd(fb))
	return fb, nil
}

func (b *FlatbufferGRPCBlober) Stream(srv service.Blober_StreamServer) error {
	ctx := srv.Context()
	req, err := srv.Recv()
	if err == io.EOF {
		return nil
	}
	defer func() {
		b := flatbuffers.NewBuilder(0)
		service.StreamResStart(b)
		b.Finish(service.StreamResEnd(b))
		srv.SendAndClose(b)
	}()

	fd, err := b.FS.Create(ctx, string(req.Name()))
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
		if _, err := fd.Write(req.BlobBytes()); err != nil {
			return status.Errorf(codes.Internal, "can't write to file: %v", err)
		}
	}
}
