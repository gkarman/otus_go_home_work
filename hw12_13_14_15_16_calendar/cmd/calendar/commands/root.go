package commands

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/infrastructure/app"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/infrastructure/config"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/infrastructure/logger"
	internalhttp "github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/infrastructure/server/http"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/infrastructure/storage"
	"github.com/spf13/cobra"
)

var (
	configPath string
	cfg        *config.Config
)

var rootCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Calendar service",
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
		var err error
		cfg, err = config.Load(configPath)
		if err != nil {
			return fmt.Errorf("ошибка загрузки конфига: %w", err)
		}
		return nil
	},
	Run: func(_ *cobra.Command, _ []string) {
		runCalendar()
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "configs/calendar_config.yaml", "Path to configuration")
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("ошибка запуска: %v", err)
	}
}

func runCalendar() {
	logg := logger.New(cfg.Logger)
	st, err := storage.New(cfg.Storage)
	if err != nil {
		logg.Error("failed to init storage: " + err.Error())
		os.Exit(1)
	}
	calendar := app.New(logg, st)
	server := internalhttp.New(cfg.Server, logg, calendar)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
	}
}
