package dbtest

import (
	"database/sql"
	"io"
	"log/slog"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/database"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBMock(t *testing.T) (*sql.DB, *database.ShoplistDB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New()
	require.NoError(t, err)

	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	require.NoError(t, err)

	db := database.NewFromGorm(gormdb)

	return sqldb, db, mock
}

func CloseWithErrorCheck(closer io.Closer) {
	if err := closer.Close(); err != nil {
		slog.Error("failed to close resource", "error", err)
	}
}
