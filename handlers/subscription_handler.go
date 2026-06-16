package handlers

import (
	"belajar-fiber/database"
	"belajar-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func CreateSubscription(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)

	var subscription models.Subscription
	if err := c.BodyParser(&subscription); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	subscription.UserID = uint(userID)
	subscription.Status = "pending"

	if err := database.DB.Create(&subscription).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create subscription",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(subscription)
}

func GetSubscriptions(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)

	var subscriptions []models.Subscription
	if err := database.DB.Where("user_id = ?", uint(userID)).Find(&subscriptions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch subscriptions",
		})
	}

	return c.JSON(subscriptions)
}

func GetAllSubscriptions(c *fiber.Ctx) error {
	var subscriptions []models.Subscription
	if err := database.DB.Find(&subscriptions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch subscriptions",
		})
	}
	return c.JSON(subscriptions)
}
