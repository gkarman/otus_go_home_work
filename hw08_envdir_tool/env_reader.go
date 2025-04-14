package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

type EnvValue struct {
	Value      string
	NeedRemove bool
}

var ErrNoValidNameFile = errors.New("no valid file name")
var ErrEntryIsDir = errors.New("entry is directory")

func ReadDir(dir string) (Environment, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		err := fmt.Errorf("read directory: %w", err)
		return nil, err
	}

	environment := make(Environment)
	for _, entry := range entries {
		key, err := getKey(entry)
		if err != nil {
			return nil, fmt.Errorf("get key: %w", err)
		}

		value, err := getValue(dir, entry)
		if err != nil {
			return nil, fmt.Errorf("get value: %w", err)
		}

		isForRemove, err := isNeedRemove(entry)
		if err != nil {
			return nil, fmt.Errorf("get is_for_remove: %w", err)
		}

		environment[key] = EnvValue{value, isForRemove}
	}

	return environment, nil
}

func getKey(entry os.DirEntry) (string, error) {
	if entry.IsDir() {
		return "", ErrEntryIsDir
	}
	fileName := entry.Name()
	badSymbols := `=`
	if strings.ContainsAny(fileName, badSymbols) {
		return "", ErrNoValidNameFile
	}
	return fileName, nil
}

func getValue(dir string, entry os.DirEntry) (string, error) {
	filePath := filepath.Join(dir, entry.Name())
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("read line: %w", err)
	}

	clean := strings.ReplaceAll(line, "\x00", "\n")
	clean = strings.TrimRight(clean, " \r\n\t%")

	return clean, nil
}

func isNeedRemove(entry os.DirEntry) (bool, error) {
	fileInfo, err := entry.Info()
	if err != nil {
		err := fmt.Errorf("get file info: %w", err)
		return false, err
	}
	if fileInfo.Size() == 0 {
		return true, nil
	}
	return false, nil
}
