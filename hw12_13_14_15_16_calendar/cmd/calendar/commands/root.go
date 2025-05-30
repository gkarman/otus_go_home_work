package commands

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/app"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/storage/memory"
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
	logg := logger.New(cfg.Logger.Level)
	storage := memorystorage.New()
	calendar := app.New(logg, storage)
	server := internalhttp.NewServer(logg, calendar)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}

	fmt.Println(cfg)
}
