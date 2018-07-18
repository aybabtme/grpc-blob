package blobersrv

import (
	"context"
	"io"

	"github.com/aybabtme/grpc-blob/blober"
	service "github.com/aybabtme/grpc-blob/gen/golanggrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ service.BloberServer = (*GolangGRPCBlober)(nil)

type GolangGRPCBlober struct {
	FS blober.Blober
}

func (b *GolangGRPCBlober) Put(ctx context.Context, req *service.PutReq) (*service.PutRes, error) {
	fd, err := b.FS.Create(ctx, req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't create file: %v", err)
	}
	if _, err := fd.Write(req.GetBlob()); err != nil {
		return nil, status.Errorf(codes.Internal, "can't write to file: %v", err)
	}
	if err := fd.Close(); err != nil {
		return nil, status.Errorf(codes.Internal, "can't close file: %v", err)
	}
	return new(service.PutRes), nil
}

func (b *GolangGRPCBlober) Stream(srv service.Blober_StreamServer) error {
	ctx := srv.Context()
	req, err := srv.Recv()
	if err == io.EOF {
		return nil
	}
	defer srv.SendAndClose(&service.StreamRes{})

	fd, err := b.FS.Create(ctx, req.GetName())
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
		if _, err := fd.Write(req.GetBlob()); err != nil {
			return status.Errorf(codes.Internal, "can't write to file: %v", err)
		}
	}
}
