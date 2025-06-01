package logger

import (
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/intrastructe/config"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/intrastructe/logger/standartlogger"
)

func New(cfg config.LoggerConf) logger.Logger {
	return standartlogger.New(cfg.Level)
}
