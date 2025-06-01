package storage

import (
	"errors"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/storage"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/intrastructe/config"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/intrastructe/storage/memorystorage"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/intrastructe/storage/sqlstorage"
)

var ErrUnknownStorageType = errors.New("неизвестный тип хранилища")

func New(cfg config.StorageConf) (storage.Storage, error) {
	switch cfg.Type {
	case "memory":
		return memorystorage.New(), nil
	case "sql":
		return sqlstorage.New(cfg), nil
	default:
		return nil, ErrUnknownStorageType
	}
}
