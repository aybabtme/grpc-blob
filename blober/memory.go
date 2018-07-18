package blober

import (
	"context"
	"io"
	"os"
	"sync"
)

type memory struct {
	blobs sync.Map
}

func Memory() Blober {
	return &memory{blobs: sync.Map{}}
}

type memoryWc struct {
	parent *memory
	name   string
	bytes  []byte
}

func (mem *memory) Write(ctx context.Context, name string, payload []byte) error {
	_, ok := mem.blobs.LoadOrStore(name, payload)
	if ok {
		return os.ErrExist
	}
	return nil
}

var nothing []byte

func (mem *memory) Create(ctx context.Context, name string) (io.WriteCloser, error) {
	_, ok := mem.blobs.LoadOrStore(name, nothing)
	if ok {
		return nil, os.ErrExist
	}
	return &memoryWc{parent: mem, name: name}, nil
}

func (w *memoryWc) Write(payload []byte) (int, error) {
	cp := make([]byte, len(payload))
	copy(cp, payload)
	w.bytes = append(w.bytes, cp...)
	return len(payload), nil
}

func (w *memoryWc) Close() error {
	w.parent.blobs.Store(w.name, w.bytes)
	return nil
}
