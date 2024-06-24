package models

type Category struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	Name         uint `json:"name"`
	Budgets      []Budget
	Transactions []Transaction
}
