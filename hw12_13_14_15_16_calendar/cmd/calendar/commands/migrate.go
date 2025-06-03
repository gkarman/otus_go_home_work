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

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Выполнить миграцию базы данных",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Запустили выполнить миграции")

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

		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("ошибка миграции: %v", err)
		}

		fmt.Println("Миграции выполнены успешно")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
