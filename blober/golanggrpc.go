package blober

import (
	"context"
	"io"
	"os"

	service "github.com/aybabtme/grpc-blob/gen/golanggrpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GolangGRPC(client service.BloberClient) Blober {
	return &golangGRPCClient{client: client}
}

type golangGRPCClient struct {
	client service.BloberClient
}

func (b *golangGRPCClient) Put(ctx context.Context, name string, blob []byte) error {
	res, err := b.client.Put(ctx, &service.PutReq{Name: name, Blob: blob})
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return os.ErrExist
		}
		return err
	}
	_ = res // not used for now
	return nil
}

func (b *golangGRPCClient) Get(ctx context.Context, name string) (blob []byte, err error) {
	res, err := b.client.Get(ctx, &service.GetReq{Name: name})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
	return res.GetBlob(), nil
}

func (c *golangGRPCClient) Write(ctx context.Context, name string) (io.WriteCloser, error) {
	srv, err := c.client.Write(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "opening stream")
	}
	req := new(service.WriteReq)
	req.Phase = &service.WriteReq_Name{Name: name}

	if err := srv.Send(req); err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, os.ErrExist
		}
		return nil, errors.Wrap(err, "sending file name")
	}
	return &golangGRPCClientWc{req: req, blob: new(service.WriteReq_Blob), srv: srv}, nil
}

type golangGRPCClientWc struct {
	req  *service.WriteReq
	blob *service.WriteReq_Blob
	srv  service.Blober_WriteClient
}

func (w *golangGRPCClientWc) Write(payload []byte) (int, error) {
	blob := w.blob
	blob.Blob = payload
	req := w.req
	req.Reset()
	req.Phase = blob
	if err := w.srv.Send(req); err != nil {
		return 0, errors.Wrap(err, "sending bytes")
	}
	return len(payload), nil
}

func (w *golangGRPCClientWc) Close() error {
	res, err := w.srv.CloseAndRecv()
	if err != nil {
		return errors.Wrap(err, "closing")
	}
	_ = res // not used for now
	return nil
}

func (c *golangGRPCClient) Read(ctx context.Context, name string, bufSize uint32) (io.ReadCloser, error) {
	srv, err := c.client.Read(ctx, &service.ReadReq{Name: name, BufSize: bufSize})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, os.ErrNotExist
		}
		return nil, errors.Wrap(err, "opening stream")
	}
	return &golangGRPCClientRc{buf: nil, srv: srv}, nil
}

type golangGRPCClientRc struct {
	buf []byte
	srv service.Blober_ReadClient
}

func (w *golangGRPCClientRc) Read(payload []byte) (int, error) {
	if len(w.buf) == 0 {
		res, err := w.srv.Recv()
		if err != nil {
			if err == io.EOF {
				return 0, err
			}
			return 0, errors.Wrap(err, "receiving bytes")
		}
		w.buf = res.GetBlob()
	}
	n := copy(payload, w.buf)
	w.buf = w.buf[n:]
	return n, nil
}

func (w *golangGRPCClientRc) Close() error {
	err := w.srv.CloseSend()
	if err != nil {
		return errors.Wrap(err, "closing")
	}
	return nil
}
