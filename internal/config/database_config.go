package config

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDbProvider(ctx context.Context) (*gorm.DB, error) {
	dsn := "host=localhost user=admin password=password dbname=personal_finance port=5432 sslmode=disable TimeZone=Asia/Kathmandu"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the database: %w", err)
	}

	return db, nil
}
