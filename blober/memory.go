package blober

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

type memory struct {
	blobs sync.Map
}

func Memory() Blober {
	return &memory{blobs: sync.Map{}}
}

func (mem *memory) Put(ctx context.Context, name string, blob []byte) error {
	mem.blobs.Store(name, blob)
	return nil
}

func (mem *memory) Get(ctx context.Context, name string) ([]byte, error) {
	blob, ok := mem.blobs.Load(name)
	if !ok {
		return nil, os.ErrNotExist
	}
	return blob.([]byte), nil
}

var nothing []byte

func (mem *memory) Write(ctx context.Context, name string) (io.WriteCloser, error) {
	_, ok := mem.blobs.LoadOrStore(name, nothing)
	if ok {
		return nil, os.ErrExist
	}
	return &memoryWc{parent: mem, name: name}, nil
}

type memoryWc struct {
	parent *memory
	name   string
	bytes  []byte
}

func (w *memoryWc) Write(blob []byte) (int, error) {
	cp := make([]byte, len(blob))
	copy(cp, blob)
	w.bytes = append(w.bytes, cp...)
	return len(blob), nil
}

func (w *memoryWc) Close() error {
	w.parent.blobs.Store(w.name, w.bytes)
	return nil
}

func (mem *memory) Read(ctx context.Context, name string) (io.ReadCloser, error) {
	blob, ok := mem.blobs.Load(name)
	if !ok {
		return nil, os.ErrNotExist
	}
	return ioutil.NopCloser(bytes.NewReader(blob.([]byte))), nil
}
