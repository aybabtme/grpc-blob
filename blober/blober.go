package blober

import (
	"context"
	"io"
)

type Blober interface {
	Put(context.Context, string, []byte) error
	Get(context.Context, string) ([]byte, error)
	Write(context.Context, string) (io.WriteCloser, error)
	Read(context.Context, string) (io.ReadCloser, error)
}
