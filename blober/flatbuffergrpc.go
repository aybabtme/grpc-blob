package blober

import (
	"context"
	"io"
	"os"

	service "github.com/aybabtme/grpc-blob/gen/flatbuffergrpc/service"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type flatbufferClient struct {
	client service.BloberClient
}

func Flatbuffer(client service.BloberClient) Blober {
	return &flatbufferClient{client: client}
}

func (c *flatbufferClient) Put(ctx context.Context, name string, payload []byte) error {
	b := flatbuffers.NewBuilder(0)
	nameT := b.CreateString(name)
	payloadT := b.CreateByteVector(payload)
	service.PutReqStart(b)
	service.PutReqAddName(b, nameT)
	service.PutReqAddBlob(b, payloadT)
	b.Finish(service.PutReqEnd(b))
	res, err := c.client.Put(ctx, b)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return os.ErrExist
		}
		return err
	}
	_ = res // not used for now
	return nil
}

func (c *flatbufferClient) Get(ctx context.Context, name string) ([]byte, error) {
	b := flatbuffers.NewBuilder(0)
	nameT := b.CreateString(name)
	service.GetReqStart(b)
	service.GetReqAddName(b, nameT)
	b.Finish(service.GetReqEnd(b))
	res, err := c.client.Get(ctx, b)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
	return res.BlobBytes(), nil
}

func (c *flatbufferClient) Write(ctx context.Context, name string) (io.WriteCloser, error) {
	srv, err := c.client.Write(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "opening stream")
	}
	b := flatbuffers.NewBuilder(0)
	nameT := b.CreateString(name)
	service.WriteReqStart(b)
	service.WriteReqAddName(b, nameT)
	b.Finish(service.WriteReqEnd(b))

	if err := srv.Send(b); err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, os.ErrExist
		}
		return nil, errors.Wrap(err, "sending file name")
	}
	return &flatbufferWc{b: b, srv: srv}, nil
}

type flatbufferWc struct {
	b   *flatbuffers.Builder
	srv service.Blober_WriteClient
}

func (w *flatbufferWc) Write(payload []byte) (int, error) {
	b := w.b
	b.Reset()
	payloadT := b.CreateByteVector(payload)
	service.WriteReqStart(b)
	service.WriteReqAddBlob(b, payloadT)
	b.Finish(service.WriteReqEnd(b))

	if err := w.srv.Send(b); err != nil {
		return 0, errors.Wrap(err, "sending bytes")
	}
	return len(payload), nil
}

func (w *flatbufferWc) Close() error {
	res, err := w.srv.CloseAndRecv()
	if err != nil {
		return errors.Wrap(err, "closing")
	}
	_ = res // not used for now
	return nil
}

func (c *flatbufferClient) Read(ctx context.Context, name string) (io.ReadCloser, error) {
	b := flatbuffers.NewBuilder(0)
	nameT := b.CreateString(name)
	service.ReadReqStart(b)
	service.ReadReqAddName(b, nameT)
	b.Finish(service.ReadReqEnd(b))
	srv, err := c.client.Read(ctx, b)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, os.ErrNotExist
		}
		return nil, errors.Wrap(err, "opening stream")
	}
	return &flatbufferRc{buf: nil, srv: srv}, nil
}

type flatbufferRc struct {
	buf []byte
	srv service.Blober_ReadClient
}

func (w *flatbufferRc) Read(payload []byte) (int, error) {
	if len(w.buf) == 0 {
		res, err := w.srv.Recv()
		if err != nil {
			return 0, errors.Wrap(err, "receiving bytes")
		}
		w.buf = res.BlobBytes()
	}
	n := copy(payload, w.buf)
	w.buf = w.buf[n:]
	return n, nil
}

func (w *flatbufferRc) Close() error {
	err := w.srv.CloseSend()
	if err != nil {
		return errors.Wrap(err, "closing")
	}
	return nil
}
