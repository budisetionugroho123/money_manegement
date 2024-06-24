package repositories

import (
	"github.com/budisetionugroho123/money_manegement/internal/models"
	"gorm.io/gorm"
)

type dbBudget struct {
	Conn *gorm.DB
}

// Create implements BudgetRepository.
func (*dbBudget) Create(budget models.Budget) (models.Budget, error) {
	panic("unimplemented")
}

// Delete implements BudgetRepository.
func (*dbBudget) Delete(id uint) error {
	panic("unimplemented")
}

// GetBudget implements BudgetRepository.
func (*dbBudget) GetBudget() ([]models.Budget, error) {
	panic("unimplemented")
}

// GetBudgetByUser implements BudgetRepository.
func (*dbBudget) GetBudgetByUser() ([]models.Budget, error) {
	panic("unimplemented")
}

// GetById implements BudgetRepository.
func (*dbBudget) GetById(id uint) (models.Budget, error) {
	panic("unimplemented")
}

// Update implements BudgetRepository.
func (*dbBudget) Update(id uint, budget models.Category) error {
	panic("unimplemented")
}

// Delete implements CategoryRepository.

type BudgetRepository interface {
	Create(budget models.Budget) (models.Budget, error)
	GetBudget() ([]models.Budget, error)
	GetById(id uint) (models.Budget, error)
	Update(id uint, budget models.Category) error
	Delete(id uint) error
	GetBudgetByUser() ([]models.Budget, error)
}

func NewBudgetRepository(Conn *gorm.DB) BudgetRepository {
	return &dbBudget{Conn: Conn}
}
