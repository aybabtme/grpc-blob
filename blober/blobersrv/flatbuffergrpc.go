package blobersrv

import (
	"context"
	"io"
	"os"

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
	err := b.FS.Put(ctx, string(req.Name()), req.BlobBytes())
	if err != nil {
		if err == os.ErrExist {
			return nil, status.Error(codes.AlreadyExists, "blob already exists")
		}
		return nil, status.Errorf(codes.Internal, "can't put blob: %v", err)
	}
	fb := flatbuffers.NewBuilder(0)
	service.PutResStart(fb)
	fb.Finish(service.PutResEnd(fb))
	return fb, nil
}

func (b *FlatbufferGRPCBlober) Get(ctx context.Context, req *service.GetReq) (*flatbuffers.Builder, error) {
	blob, err := b.FS.Get(ctx, string(req.Name()))
	if err != nil {
		if err == os.ErrNotExist {
			return nil, status.Error(codes.NotFound, "blob not found")
		}
		return nil, status.Errorf(codes.Internal, "can't get blob: %v", err)
	}
	fb := flatbuffers.NewBuilder(0)
	blobT := fb.CreateByteVector(blob)
	service.GetResStart(fb)
	service.GetResAddBlob(fb, blobT)
	fb.Finish(service.GetResEnd(fb))
	return fb, nil
}

func (b *FlatbufferGRPCBlober) Write(srv service.Blober_WriteServer) error {
	ctx := srv.Context()
	req, err := srv.Recv()
	if err == io.EOF {
		return nil
	}
	defer func() {
		b := flatbuffers.NewBuilder(0)
		service.WriteResStart(b)
		b.Finish(service.WriteResEnd(b))
		srv.SendAndClose(b)
	}()

	fd, err := b.FS.Write(ctx, string(req.Name()))
	if err != nil {
		if err == os.ErrExist {
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
		if _, err := fd.Write(req.BlobBytes()); err != nil {
			return status.Errorf(codes.Internal, "can't write to blob: %v", err)
		}
	}
}

func (b *FlatbufferGRPCBlober) Read(req *service.ReadReq, srv service.Blober_ReadServer) error {
	ctx := srv.Context()

	r, err := b.FS.Read(ctx, string(req.Name()))
	if err != nil {
		if err == os.ErrNotExist {
			return status.Errorf(codes.NotFound, "blob not found")
		}
		return status.Errorf(codes.Internal, "can't open blob: %v", err)
	}
	defer r.Close()

	fb := flatbuffers.NewBuilder(0)
	buf := make([]byte, 1<<16) // 64 KiB
	for {
		n, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return status.Errorf(codes.Internal, "can't read blob: %v", err)
		}
		fb.Reset()
		blobT := fb.CreateByteVector(buf[:n])
		service.ReadResStart(fb)
		service.ReadResAddBlob(fb, blobT)
		fb.Finish(service.ReadResEnd(fb))

		if err := srv.Send(fb); err != nil {
			return status.Errorf(codes.Internal, "can't send: %v", err)
		}
	}
}
