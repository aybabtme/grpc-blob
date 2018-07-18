package blober

import (
	"context"
)

type WriteCloser interface {
	Write(context.Context, []byte) (int, error)
	Close(context.Context) error
}

type Blober interface {
	Create(context.Context, string) (WriteCloser, error)
	Write(context.Context, string, []byte) error
}
