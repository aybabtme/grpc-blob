package blober

import (
	"context"
	"io"
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

func (fs *filesystem) Create(ctx context.Context, name string) (io.WriteCloser, error) {
	path := filepath.Join(fs.root, name)
	if strings.Contains(name, "..") {
		return nil, errors.New("name may not contain `..`")
	}
	fd, err := os.Create(path)
	if err != nil {
		return nil, errors.Wrap(err, "creating file")
	}
	return fd, nil
}
