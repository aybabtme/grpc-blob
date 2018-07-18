package blober

import "io"

type Blober interface {
	Create(string) (io.WriteCloser, error)
}
