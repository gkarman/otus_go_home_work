package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

var ErrConfigNotFound = errors.New("файл конфигурации не найден")
var ErrConfigUnmarshal = errors.New("ошибка разбора конфига")

type Config struct {
	Logger  LoggerConf  `yaml:"logger"`
	Storage StorageConf `yaml:"storage"`
}

type LoggerConf struct {
	Level string `yaml:"level"`
}

type StorageConf struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
}

type ServerConf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func Load(f string) (*Config, error) {
	data, err := os.ReadFile(f)
	if err != nil {
		return nil, ErrConfigNotFound
	}

	var c Config
	if err := yaml.Unmarshal(data, &c); err != nil {
		return nil, ErrConfigUnmarshal
	}

	return &c, nil
}
