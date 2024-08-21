package database

import (
	"fmt"

	"bff-todolist/internal/domain"
	"bff-todolist/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.Todo{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
