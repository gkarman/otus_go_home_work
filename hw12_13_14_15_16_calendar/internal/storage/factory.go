package storage

import (
	"errors"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/storage/memorystorage"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/storage/sqlstorage"
)

var ErrUnknownStorageType = errors.New("неизвестный тип хранилища")

func NewStorage(cfg config.StorageConf) (StorageInterface, error) {
	switch cfg.Type {
	case "memory":
		return memorystorage.New(), nil
	case "sql":
		return sqlstorage.New(), nil
	default:
		return nil, ErrUnknownStorageType
	}
}
