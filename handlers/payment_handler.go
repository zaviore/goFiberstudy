package handlers

import (
	"belajar-fiber/database"
	"belajar-fiber/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PaymentRequest struct {
	SubscriptionID uint   `json:"subscription_id"`
	Method         string `json:"method"`
}

func ProcessPayment(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)

	var req PaymentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var subscription models.Subscription
	if err := database.DB.Where("id = ? AND user_id = ?", req.SubscriptionID, uint(userID)).First(&subscription).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Subscription not found",
		})
	}

	if subscription.Status == "paid" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Subscription already paid",
		})
	}

	payment := models.Payment{
		SubscriptionID: req.SubscriptionID,
		Method:         req.Method,
		Status:         "pending",
		TransactionID:  uuid.New().String(),
	}

	if err := database.DB.Create(&payment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create payment",
		})
	}

	go func() {
		time.Sleep(2 * time.Second)
		database.DB.Model(&payment).Update("status", "completed")
		database.DB.Model(&subscription).Update("status", "paid")
	}()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":        "Payment processed successfully",
		"payment":        payment,
		"transaction_id": payment.TransactionID,
	})
}

func GetPayments(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)

	var payments []models.Payment
	if err := database.DB.Joins("JOIN subscriptions ON payments.subscription_id = subscriptions.id").
		Where("subscriptions.user_id = ?", uint(userID)).
		Find(&payments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch payments",
		})
	}

	return c.JSON(payments)
}

func GetPaymentStatus(c *fiber.Ctx) error {
	transactionID := c.Params("transaction_id")

	var payment models.Payment
	if err := database.DB.Where("transaction_id = ?", transactionID).First(&payment).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Payment not found",
		})
	}

	return c.JSON(fiber.Map{
		"transaction_id": payment.TransactionID,
		"status":         payment.Status,
		"method":         payment.Method,
	})
}
