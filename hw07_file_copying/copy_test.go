package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("err file not exist", func(t *testing.T) {
		pathFrom := "./file"
		pathTo := "./file2"
		offset := int64(0)
		limit := int64(0)

		t.Cleanup(func() {
			_ = os.Remove(pathTo)
		})

		err := Copy(pathFrom, pathTo, offset, limit)
		require.Error(t, err)
		require.True(t, errors.Is(err, os.ErrNotExist), "Ожидалась ошибка os.ErrNotExist, но получили: %v", err)
	})

	t.Run("file has not size", func(t *testing.T) {
		pathFrom := "/dev/urandom"
		pathTo := "./file2"
		offset := int64(0)
		limit := int64(0)

		t.Cleanup(func() {
			_ = os.Remove(pathTo)
		})

		err := Copy(pathFrom, pathTo, offset, limit)
		require.Error(t, err)
		require.True(t, errors.Is(err, ErrUnsupportedFile), "Ожидалась ErrUnsupportedFile, но получили: %v", err)
	})

	t.Run("no valid offset", func(t *testing.T) {
		pathFrom := "testdata/input.txt"
		pathTo := "temp.txt"
		limit := int64(0)

		fileFromInfo, err := os.Stat(pathFrom)
		if err != nil {
			t.Fatal(err)
		}

		fileFromSize := fileFromInfo.Size()
		offset := fileFromSize + 1

		t.Cleanup(func() {
			_ = os.Remove(pathTo)
		})

		err = Copy(pathFrom, pathTo, offset, limit)
		require.Error(t, err)
		require.True(t, errors.Is(err, ErrOffsetExceedsFileSize), "Ожидалась ErrOffsetExceedsFileSize, но получили: %v", err)
	})
}
