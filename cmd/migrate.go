package cmd

import (
	"errors"
	"log/slog"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // required for postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // required for file source
	"github.com/spf13/cobra"

	"github.com/cristiano-pacheco/goflix/internal/shared/modules/config"
	"github.com/cristiano-pacheco/goflix/pkg/database"
)

// dbMigrateCmd represents the migrate command.
var dbMigrateCmd = &cobra.Command{
	Use:   "db:migrate",
	Short: "Run database migrations",
	Long:  `Run database migrations. This command will run all the migrations that have not been run yet.`,
	Run: func(_ *cobra.Command, _ []string) {
		config.Init()
		cfg := config.GetConfig()
		dbConfig := database.Config{
			Host:               cfg.DB.Host,
			User:               cfg.DB.User,
			Password:           cfg.DB.Password,
			Name:               cfg.DB.Name,
			Port:               cfg.DB.Port,
			MaxOpenConnections: cfg.DB.MaxOpenConnections,
			MaxIdleConnections: cfg.DB.MaxIdleConnections,
			SSLMode:            cfg.DB.SSLMode,
			PrepareSTMT:        cfg.DB.PrepareSTMT,
			EnableLogs:         cfg.DB.EnableLogs,
		}
		dsn := database.GeneratePostgresDatabaseDSN(dbConfig)

		m, err := migrate.New("file://migrations", dsn)
		if err != nil {
			slog.Error("Failed to create migration instance", "error", err)
			os.Exit(1)
		}

		err = m.Up()
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			slog.Error("Failed to run migrations", "error", err)
			os.Exit(1)
		}

		slog.Info("Migrations executed successfully")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(dbMigrateCmd)
}
