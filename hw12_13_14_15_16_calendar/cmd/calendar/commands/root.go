package commands

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/intrastructe/app"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/intrastructe/config"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/intrastructe/logger"
	internalhttp "github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/intrastructe/server/http"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/intrastructe/storage"
	"github.com/spf13/cobra"
)

var (
	configPath string
	cfg        *config.Config
)

var rootCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Calendar service",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		cfg, err = config.Load(configPath)
		if err != nil {
			return fmt.Errorf("ошибка загрузки конфига: %w", err)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		runCalendar()
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "configs/calendar_config.yaml", "Path to configuration file")
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("ошибка запуска: %v", err)
	}
}

func runCalendar() {
	logg := logger.New(cfg.Logger)
	st, err := storage.New(cfg.Storage)
	if err != nil {
		logg.Error("failed to init storage: " + err.Error())
		os.Exit(1) //nolint:gocritic
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
		os.Exit(1) //nolint:gocritic
	}
}
