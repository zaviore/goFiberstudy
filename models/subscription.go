package models

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model
	UserID    uint   `json:"user_id" gorm:"not null"`
	PackageID string `json:"package_id" gorm:"not null"`
	Status    string `json:"status" gorm:"default:pending"`
	Amount    float64 `json:"amount"`
}

type Payment struct {
	gorm.Model
	SubscriptionID uint   `json:"subscription_id" gorm:"not null"`
	Method         string `json:"method"`
	Status         string `json:"status" gorm:"default:pending"`
	TransactionID  string `json:"transaction_id"`
}
