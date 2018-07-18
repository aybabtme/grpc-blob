package blober

import (
	"context"

	service "github.com/aybabtme/grpc-blob/gen/flatbuffer/service"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/pkg/errors"
)

type flatbufferClient struct {
	client service.BloberClient
}

func Flatbuffer(client service.BloberClient) Blober {
	return &flatbufferClient{client: client}
}

func (c *flatbufferClient) Write(ctx context.Context, name string, payload []byte) error {
	b := flatbuffers.NewBuilder(0)
	nameT := b.CreateString(name)
	payloadT := b.CreateByteVector(payload)
	service.StreamReqStart(b)
	service.StreamReqAddName(b, nameT)
	service.StreamReqAddBlob(b, payloadT)
	b.Finish(service.StreamReqEnd(b))
	res, err := c.client.Put(ctx, b)
	if err != nil {
		return errors.Wrap(err, "writing")
	}
	_ = res // not used for now
	return nil
}

type flatbufferWc struct {
	b   *flatbuffers.Builder
	srv service.Blober_StreamClient
}

func (c *flatbufferClient) Create(ctx context.Context, name string) (WriteCloser, error) {
	srv, err := c.client.Stream(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "opening stream")
	}
	b := flatbuffers.NewBuilder(0)
	nameT := b.CreateString(name)
	service.StreamReqStart(b)
	service.StreamReqAddName(b, nameT)
	b.Finish(service.StreamReqEnd(b))

	if err := srv.Send(b); err != nil {
		return nil, errors.Wrap(err, "sending file name")
	}
	return &flatbufferWc{b: b, srv: srv}, nil
}

func (w *flatbufferWc) Write(ctx context.Context, payload []byte) (int, error) {
	b := w.b
	b.Reset()
	payloadT := b.CreateByteVector(payload)
	service.StreamReqStart(b)
	service.StreamReqAddBlob(b, payloadT)
	b.Finish(service.StreamReqEnd(b))

	if err := w.srv.Send(b); err != nil {
		return 0, errors.Wrap(err, "sending bytes")
	}
	return len(payload), nil
}

func (w *flatbufferWc) Close(ctx context.Context) error {
	res, err := w.srv.CloseAndRecv()
	if err != nil {
		return errors.Wrap(err, "closing")
	}
	_ = res // not used for now
	return nil
}
