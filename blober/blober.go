package blober

import (
	"context"
	"io"
)

type Blober interface {
	Create(context.Context, string) (io.WriteCloser, error)
	Write(context.Context, string, []byte) error
}
