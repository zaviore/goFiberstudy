package database

import (
	"belajar-fiber/config"
	"belajar-fiber/models"
	"fmt"
	"log"
)

func RunMigrations(cfg *config.Config) error {
	log.Println("Starting database migrations...")

	if DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	log.Println("Auto-migrating models...")
	
	err := DB.AutoMigrate(
		&models.User{},
		&models.Subscription{},
		&models.Payment{},
	)
	if err != nil {
		return fmt.Errorf("failed to auto-migrate: %w", err)
	}

	log.Println("Database migrations completed successfully!")
	return nil
}

func DropTables() error {
	log.Println("Dropping all tables...")
	
	err := DB.Migrator().DropTable(
		&models.Payment{},
		&models.Subscription{},
		&models.User{},
	)
	if err != nil {
		return fmt.Errorf("failed to drop tables: %w", err)
	}

	log.Println("All tables dropped successfully!")
	return nil
}

func ResetDatabase() error {
	log.Println("Resetting database...")
	
	if err := DropTables(); err != nil {
		return err
	}
	
	if err := RunMigrations(config.AppCfg); err != nil {
		return err
	}

	log.Println("Database reset completed successfully!")
	return nil
}
