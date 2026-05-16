package database

import (
	"fmt"
	"github.com/noctivs/smart-community-server/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitPostgres() {
	cfg := utils.Config

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.GetString("postgres.host"),
		cfg.GetString("postgres.user"),
		cfg.GetString("postgres.password"),
		cfg.GetString("postgres.database"),
		cfg.GetInt("postgres.port"),
		cfg.GetString("postgres.sslmode"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
