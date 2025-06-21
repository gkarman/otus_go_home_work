package commands

import (
	"fmt"
	"log"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/infrastructure/config"
	"github.com/spf13/cobra"
)

var (
	configPath string
	cfg        *config.Config
)

var rootCmd = &cobra.Command{
	Use:   "calendar scheduler",
	Short: "Calendar scheduler service",
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
		var err error
		cfg, err = config.Load(configPath)
		if err != nil {
			return fmt.Errorf("ошибка загрузки конфига: %w", err)
		}
		if cfg.Broker == nil || cfg.Broker.Type != "rabbitmq" {
			log.Fatal("RabbitMQ config is required")
		}
		return nil
	},
	Run: func(_ *cobra.Command, _ []string) {
		runCalendarScheduler()
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "configs/scheduler_config.yaml", "Path to configuration")
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("ошибка запуска: %v", err)
	}
}

func runCalendarScheduler() {
	fmt.Println("hello")
	fmt.Println(cfg.Broker)

	producer, err := rabbitmq.NewRabbitProducer(cfg.Broker)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

}
