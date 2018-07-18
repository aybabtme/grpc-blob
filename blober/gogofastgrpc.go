package blober

import (
	"context"
	"io"

	service "github.com/aybabtme/grpc-blob/gen/gogofastgrpc"
	"github.com/pkg/errors"
)

func GogoFastGRPC(client service.BloberClient) Blober {
	return &gogoFastGRPCClient{client: client}
}

type gogoFastGRPCClient struct {
	client service.BloberClient
}

func (b *gogoFastGRPCClient) Write(ctx context.Context, name string, payload []byte) error {
	return nil
}

type gogoFastGRPCClientWc struct {
	req  *service.StreamReq
	blob *service.StreamReq_Blob
	srv  service.Blober_StreamClient
}

func (c *gogoFastGRPCClient) Create(ctx context.Context, name string) (io.WriteCloser, error) {
	srv, err := c.client.Stream(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "opening stream")
	}
	req := new(service.StreamReq)
	req.Phase = &service.StreamReq_Name{Name: name}

	if err := srv.Send(req); err != nil {
		return nil, errors.Wrap(err, "sending file name")
	}
	return &gogoFastGRPCClientWc{req: req, blob: new(service.StreamReq_Blob), srv: srv}, nil
}

func (w *gogoFastGRPCClientWc) Write(payload []byte) (int, error) {
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

func (w *gogoFastGRPCClientWc) Close() error {
	res, err := w.srv.CloseAndRecv()
	if err != nil {
		return errors.Wrap(err, "closing")
	}
	_ = res // not used for now
	return nil
}
