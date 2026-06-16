package main

import (
	"belajar-fiber/config"
	"belajar-fiber/database"
	"belajar-fiber/routes"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	envFlag := flag.String("env", "development", "Environment to use (development/staging/production)")
	migrateFlag := flag.Bool("migrate", false, "Run database migrations")
	resetFlag := flag.Bool("reset", false, "Reset database (drop all tables and migrate)")
	flag.Parse()

	cfg, err := config.LoadConfig(*envFlag)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := database.ConnectDB(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if *resetFlag {
		if err := database.ResetDatabase(); err != nil {
			log.Fatalf("Failed to reset database: %v", err)
		}
		return
	}

	if *migrateFlag {
		if err := database.RunMigrations(cfg); err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		return
	}

	if err := database.RunMigrations(cfg); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	routes.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Subscription API",
			"version": "1.0.0",
			"env": cfg.App.Environment,
		})
	})

	log.Printf("Starting server on port %s in %s mode...", cfg.App.Port, cfg.App.Environment)
	app.Listen(":" + cfg.App.Port)
}
