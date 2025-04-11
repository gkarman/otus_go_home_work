package main

import (
	"bufio"
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

func ReadDir(dir string) (Environment, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		err := fmt.Errorf("Ошибка чтения директории")
		return nil, err
	}

	environment := make(Environment)
	for _, entry := range entries {
		key, err := getKey(entry)
		if err != nil {
			return nil, err
		}

		value, err := getValue(dir, entry)
		if err != nil {
			return nil, err
		}

		isForRemove, err := isNeedRemove(entry)
		if err != nil {
			return nil, err
		}

		environment[key] = EnvValue{value, isForRemove}
	}

	return environment, nil
}

func getKey(entry os.DirEntry) (string, error) {
	if entry.IsDir() {
		err := fmt.Errorf("Директория, нужен файл")
		return "", err
	}
	fileName := entry.Name()
	badSymbols := `=`
	if strings.ContainsAny(fileName, badSymbols) {
		err := fmt.Errorf("Имя файла не валидно")
		return "", err
	}
	return fileName, nil
}

func getValue(dir string, entry os.DirEntry) (string, error) {
	filePath := filepath.Join(dir, entry.Name())
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}

	clean := strings.ReplaceAll(line, "\x00", "\n")
	clean = strings.TrimRight(clean, " \r\n\t%")

	return clean, nil
}

func isNeedRemove(entry os.DirEntry) (bool, error) {
	fileInfo, err := entry.Info()
	if err != nil {
		err := fmt.Errorf("Ошибка чтения инфромации о файле")
		return false, err
	}
	if fileInfo.Size() == 0 {
		return true, nil
	}
	return false, nil
}
