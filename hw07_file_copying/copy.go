package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fileFromInfo, err := os.Stat(fromPath)
	if err != nil {
		return err
	}

	fileSize := fileFromInfo.Size()
	if fileSize == 0 {
		return ErrUnsupportedFile
	}

	if limit == 0 {
		limit = fileSize
	}

	fileFrom, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer fileFrom.Close()

	if offset > fileFromInfo.Size() {
		return ErrOffsetExceedsFileSize
	}

	if offset > 0 {
		if _, err := fileFrom.Seek(offset, io.SeekStart); err != nil {
			return err
		}
	}

	fileTo, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer fileTo.Close()

	progressBar := pb.Full.Start64(limit)
	fileFromPb := progressBar.NewProxyReader(fileFrom)

	_, err = io.CopyN(fileTo, fileFromPb, limit)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}
	progressBar.Finish()

	return nil
}
