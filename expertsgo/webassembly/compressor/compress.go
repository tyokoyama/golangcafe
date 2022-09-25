package compressor

import (
	"bytes"
	"compress/gzip"
	"io"
	"time"
)

func Compress(src io.Reader) (io.Reader, error) {
	var buf bytes.Buffer

	zw := gzip.NewWriter(&buf)
	zw.ModTime = time.Now()

	if _, err := io.Copy(zw, src); err != nil {
		return nil, err
	}

	if err := zw.Close(); err != nil {
		return nil, err
	}

	return &buf, nil
}