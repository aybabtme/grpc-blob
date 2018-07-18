package blober

import (
	"context"
	"io"
	"os"

	service "github.com/aybabtme/grpc-blob/gen/gofastgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ service.BloberServer = (*GoFastGRPCBlober)(nil)

type GoFastGRPCBlober struct {
	FS Blober
}

func (b *GoFastGRPCBlober) Put(ctx context.Context, req *service.PutReq) (*service.PutRes, error) {
	fd, err := b.FS.Create(req.GetName())
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

func (b *GoFastGRPCBlober) Stream(srv service.Blober_StreamServer) error {
	req, err := srv.Recv()
	if err == io.EOF {
		return nil
	}
	defer srv.SendAndClose(&service.StreamRes{})

	fd, err := b.FS.Create(req.GetName())
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
		if _, err := os.Stdout.Write(req.GetBlob()); err != nil {
			return status.Errorf(codes.Internal, "can't write to file: %v", err)
		}
	}
}
