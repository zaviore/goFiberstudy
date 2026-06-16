package database

import (
	"belajar-fiber/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) error {
	dsn := config.GetDSN(cfg.Database)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Printf("Database connected successfully to %s@%s:%s/%s", 
		cfg.Database.User, 
		cfg.Database.Host, 
		cfg.Database.Port, 
		cfg.Database.Name,
	)

	DB = db
	return nil
}
