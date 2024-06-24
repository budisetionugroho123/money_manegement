package repositories

import (
	"github.com/budisetionugroho123/money_manegement/internal/models"
	"gorm.io/gorm"
)

type dbCategory struct {
	Conn *gorm.DB
}

// Delete implements CategoryRepository.
func (db *dbCategory) Delete(id uint) error {
	var category models.Category
	return db.Conn.Where("id", id).Delete(&category).Error
}

// GetById implements CategoryRepository.
func (db *dbCategory) GetById(id uint) (models.Category, error) {
	var category models.Category
	result := db.Conn.Where("id", id).First(&category)
	return category, result.Error
}

// GetCategory implements CategoryRepository.
func (*dbCategory) GetCategory() ([]models.Category, error) {
	panic("unimplemented")
}

// Update implements CategoryRepository.
func (*dbCategory) Update(id uint, category models.Category) error {
	panic("unimplemented")
}

// Create implements CategoryRepository.
func (*dbCategory) Create(category models.Category) (models.Category, error) {
	panic("unimplemented")
}

type CategoryRepository interface {
	Create(category models.Category) (models.Category, error)
	GetCategory() ([]models.Category, error)
	GetById(id uint) (models.Category, error)
	Update(id uint, category models.Category) error
	Delete(id uint) error
}

func NewCategoryRepository(Conn *gorm.DB) CategoryRepository {
	return &dbCategory{Conn: Conn}
}
