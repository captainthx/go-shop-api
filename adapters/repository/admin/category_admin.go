package adminRepository

import (
	"go-shop-api/core/domain"
	adminPorts "go-shop-api/core/ports/admin"

	"gorm.io/gorm"
)

type categoryAdminRepositoryDB struct {
	db *gorm.DB
}

func NewCategoryAdminRepositoryDB(db *gorm.DB) adminPorts.CategoryAdminRepository {
	return &categoryAdminRepositoryDB{db: db}
}

// CreateCategory implements adminPorts.CategoryAdminRepository.
func (c *categoryAdminRepositoryDB) CreateCategory(category *domain.Category) error {
	err := c.db.Create(category).Error
	if err != nil {
		return err
	}
	return nil
}

// FindAllCategory implements adminPorts.CategoryAdminRepository.
func (c *categoryAdminRepositoryDB) FindAllCategory() ([]domain.Category, error) {
	var category []domain.Category
	resutl := c.db.Find(&category)
	if resutl.Error != nil {
		return nil, resutl.Error
	}

	return category, nil
}
