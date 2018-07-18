package blober

import (
	"context"
	"io"

	service "github.com/aybabtme/grpc-blob/gen/gofastgrpc"
	"github.com/pkg/errors"
)

func GoFastGRPC(client service.BloberClient) Blober {
	return &goFastGRPCClient{client: client}
}

type goFastGRPCClient struct {
	client service.BloberClient
}

func (b *goFastGRPCClient) Write(ctx context.Context, name string, payload []byte) error {
	return nil
}

type goFastGRPCClientWc struct {
	req  *service.StreamReq
	blob *service.StreamReq_Blob
	srv  service.Blober_StreamClient
}

func (c *goFastGRPCClient) Create(ctx context.Context, name string) (io.WriteCloser, error) {
	srv, err := c.client.Stream(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "opening stream")
	}
	req := new(service.StreamReq)
	req.Phase = &service.StreamReq_Name{Name: name}

	if err := srv.Send(req); err != nil {
		return nil, errors.Wrap(err, "sending file name")
	}
	return &goFastGRPCClientWc{req: req, blob: new(service.StreamReq_Blob), srv: srv}, nil
}

func (w *goFastGRPCClientWc) Write(payload []byte) (int, error) {
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

func (w *goFastGRPCClientWc) Close() error {
	res, err := w.srv.CloseAndRecv()
	if err != nil {
		return errors.Wrap(err, "closing")
	}
	_ = res // not used for now
	return nil
}
