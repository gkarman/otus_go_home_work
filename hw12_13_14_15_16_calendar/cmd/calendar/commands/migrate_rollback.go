package commands

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	// так нужно.
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	// так нужно.
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migrateRollbackCmd = &cobra.Command{
	Use:   "migrate:rollback",
	Short: "Откатить последнюю миграцию",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Откат последней миграции...")

		dsn := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.Storage.User,
			cfg.Storage.Password,
			cfg.Storage.Host,
			cfg.Storage.Port,
			cfg.Storage.DB,
		)

		m, err := migrate.New("file://migrations", dsn)
		if err != nil {
			log.Fatalf("ошибка создания мигратора: %v", err)
		}

		if err := m.Steps(-1); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("ошибка при откате миграции: %v", err)
		}

		fmt.Println("Миграция успешно откатилась на 1 шаг")
	},
}

func init() {
	rootCmd.AddCommand(migrateRollbackCmd)
}
