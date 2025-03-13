package utils

import (
	"bufio"
	"compress/gzip"
	"io"
	"os"
)

func CompressFile(src, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	read := bufio.NewReader(f)
	data, err := io.ReadAll(read)
	if err != nil {
		return err
	}

	gzipped, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer gzipped.Close()

	w := gzip.NewWriter(gzipped)
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	defer w.Close()

	return nil
}
