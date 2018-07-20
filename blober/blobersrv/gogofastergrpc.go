package blobersrv

import (
	"context"
	"io"
	"os"

	"github.com/aybabtme/grpc-blob/blober"
	service "github.com/aybabtme/grpc-blob/gen/gogofastergrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ service.BloberServer = (*GogoFasterGRPCBlober)(nil)

type GogoFasterGRPCBlober struct {
	FS blober.Blober
}

func (b *GogoFasterGRPCBlober) Put(ctx context.Context, req *service.PutReq) (*service.PutRes, error) {
	err := b.FS.Put(ctx, req.GetName(), req.GetBlob())
	if err != nil {
		if os.IsExist(err) {
			return nil, status.Errorf(codes.AlreadyExists, "blob already exists")
		}
		return nil, status.Errorf(codes.Internal, "can't create blob: %v", err)
	}
	return new(service.PutRes), nil
}

func (b *GogoFasterGRPCBlober) Get(ctx context.Context, req *service.GetReq) (*service.GetRes, error) {
	blob, err := b.FS.Get(ctx, req.GetName())
	if err != nil {
		if os.IsNotExist(err) {
			return nil, status.Error(codes.NotFound, "blob not found")
		}
		return nil, status.Errorf(codes.Internal, "can't get blob: %v", err)
	}
	return &service.GetRes{Blob: blob}, nil
}

func (b *GogoFasterGRPCBlober) Write(srv service.Blober_WriteServer) error {
	ctx := srv.Context()
	req, err := srv.Recv()
	if err == io.EOF {
		return nil
	}
	defer srv.SendAndClose(&service.WriteRes{})

	fd, err := b.FS.Write(ctx, req.GetName())
	if err != nil {
		if os.IsExist(err) {
			return status.Error(codes.AlreadyExists, "blob already exists")
		}
		return status.Errorf(codes.Internal, "can't create blob: %v", err)
	}
	defer fd.Close() // in any case
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			if err := fd.Close(); err != nil {
				return status.Errorf(codes.Internal, "can't close blob: %v", err)
			}
			return nil
		}
		if _, err := fd.Write(req.GetBlob()); err != nil {
			return status.Errorf(codes.Internal, "can't write to blob: %v", err)
		}
	}
}

func (b *GogoFasterGRPCBlober) Read(req *service.ReadReq, srv service.Blober_ReadServer) error {
	ctx := srv.Context()

	fd, err := b.FS.Read(ctx, req.GetName())
	if err != nil {
		if os.IsNotExist(err) {
			return status.Errorf(codes.NotFound, "blob not found")
		}
		return status.Errorf(codes.Internal, "can't create blob: %v", err)
	}
	defer fd.Close() // in any case

	res := &service.ReadRes{}
	buf := make([]byte, 1<<16) // 64 KiB
	for {
		n, err := fd.Read(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return status.Errorf(codes.Internal, "can't open blob: %v", err)
		}
		res.Reset()
		res.Blob = buf[:n]
		if err := srv.Send(res); err != nil {
			return status.Errorf(codes.Internal, "can't send: %v", err)
		}
	}
}
