package database

import (
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/config"
	"github.com/cristiano-pacheco/goflix/pkg/database"
	"gorm.io/gorm"
)

type ShoplistDB struct {
	*gorm.DB
}

func New(cfg config.Config) *ShoplistDB {
	dbConfig := database.DatabaseConfig{
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

	db := database.OpenConnection(dbConfig)

	return &ShoplistDB{DB: db}
}

func NewFromGorm(db *gorm.DB) *ShoplistDB {
	return &ShoplistDB{db}
}
