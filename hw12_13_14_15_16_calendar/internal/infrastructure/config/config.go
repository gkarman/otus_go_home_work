package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	ErrConfigNotFound  = errors.New("файл конфигурации не найден")
	ErrConfigUnmarshal = errors.New("ошибка разбора конфига")
)

type Config struct {
	Logger     LoggerConf     `yaml:"logger"`
	Storage    StorageConf    `yaml:"storage"`
	Server     ServerConf     `yaml:"server"`
	ServerGrpc ServerGrpcConf `yaml:"serverGrpc"`
	Broker     *BrokerConf    `yaml:"broker,omitempty"`
}

type LoggerConf struct {
	Level         string `yaml:"level"`
	PathToHTTPLog string `yaml:"pathToHttpLog"`
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

type ServerGrpcConf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type BrokerConf struct {
	Type        string `yaml:"type"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	VirtualHost string `yaml:"virtual_host"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Queue       string `yaml:"queue"`
	Exchange    string `yaml:"exchange"`
	RoutingKey  string `yaml:"routing_key"`
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
