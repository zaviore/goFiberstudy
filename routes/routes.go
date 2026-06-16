package routes

import (
	"belajar-fiber/handlers"
	"belajar-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	users := api.Group("/users")
	users.Get("/", handlers.GetUsers)
	users.Get("/:id", handlers.GetUser)
	users.Put("/:id", handlers.UpdateUser)
	users.Delete("/:id", handlers.DeleteUser)

	subscriptions := api.Group("/subscriptions")
	subscriptions.Use(middleware.AuthMiddleware)
	subscriptions.Post("/", handlers.CreateSubscription)
	subscriptions.Get("/", handlers.GetSubscriptions)
	subscriptions.Get("/all", handlers.GetAllSubscriptions)

	payments := api.Group("/payments")
	payments.Use(middleware.AuthMiddleware)
	payments.Post("/", handlers.ProcessPayment)
	payments.Get("/", handlers.GetPayments)
	payments.Get("/:transaction_id", handlers.GetPaymentStatus)
}
