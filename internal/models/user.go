package models

import "time"

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	Password     string `json:"password"`
	Budgets      []Budget
	Transactions []Transaction
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
