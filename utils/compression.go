package utils

import (
	"bufio"
	"compress/gzip"
	"io"
	"os"
	"strings"
)

type Uncompressor interface {
	Uncompress(filepath string) (string, error)
}

type GzipUncompressor struct{}

func (GzipUncompressor) Uncompress(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}
	defer gzipReader.Close()

	buffer, err := io.ReadAll(gzipReader)
	if err != nil {
		return "", err
	}

	outpath := strings.TrimSuffix(filepath, ".gz")

	outFile, err := os.Create(outpath)
	if err != nil {
		return "", err
	}

	defer outFile.Close()

	w := bufio.NewWriter(outFile)
	_, err = w.Write(buffer)
	if err != nil {
		return "", err
	}

	defer w.Flush()

	return outpath, nil
}

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
