package models

import (
	"time"
)

type Transaction struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Type           string    `json:"type"`
	Amount         float64   `json:"amount"`
	Description    string    `json:"description" gorm:"type:text"`
	TrasactionDate time.Time `json:"transaction_date"`
	CategoryID     uint      `json:"category_id"`
	UserID         uint      `json:"user_id"`
	BudgetID       uint      `json:"budget_id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
