package blober

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

type filesystem struct {
	root string
}

func FileSystem(root string) Blober {
	return &filesystem{root: root}
}

func (fs *filesystem) Write(ctx context.Context, name string, payload []byte) error {
	path := filepath.Join(fs.root, name)
	if strings.Contains(name, "..") {
		return errors.New("name may not contain `..`")
	}
	fd, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "creating file")
	}
	if _, err := fd.Write(payload); err != nil {
		return errors.Wrap(err, "writing payload")
	}
	if err := fd.Close(); err != nil {
		return errors.Wrap(err, "closing file")
	}
	return nil
}

type filesystemWc struct {
	fd *os.File
}

func (fs *filesystem) Create(ctx context.Context, name string) (WriteCloser, error) {
	path := filepath.Join(fs.root, name)
	if strings.Contains(name, "..") {
		return nil, errors.New("name may not contain `..`")
	}
	fd, err := os.Create(path)
	if err != nil {
		return nil, errors.Wrap(err, "creating file")
	}
	return &filesystemWc{fd: fd}, nil
}

func (w *filesystemWc) Write(ctx context.Context, payload []byte) (int, error) {
	return w.fd.Write(payload)
}

func (w *filesystemWc) Close(ctx context.Context) error {
	return w.fd.Close()
}
