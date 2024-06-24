package models

import "time"

type Budget struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	Category     string  `json:"category"`
	BudgetAmount float64 `json:"budget_amount"`
	CategoryId   uint    `json:"category_id"`
	UserId       uint    `json:"user_id"`
	Transactions []Transaction
	StartDate    time.Time
	EndDate      time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
