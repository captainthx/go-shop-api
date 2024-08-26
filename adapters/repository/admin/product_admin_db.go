package adminRepository

import (
	"go-shop-api/core/domain"
	adminPorts "go-shop-api/core/ports/admin"

	"gorm.io/gorm"
)

type productAdminRepositoryDB struct {
	db *gorm.DB
}

func NewProductAdminRepositoryDB(db *gorm.DB) adminPorts.ProductAdminRepository {
	return &productAdminRepositoryDB{db: db}
}

// CreateProduct implements adminPorts.ProductAdminRepository.
func (p *productAdminRepositoryDB) CreateProduct(product *domain.Product) error {
	if err := p.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

// DeleteProduct implements adminPorts.ProductAdminRepository.
func (p *productAdminRepositoryDB) DeleteProduct(product *domain.Product) error {
	panic("unimplemented")
}

// FindCategoryByID implements adminPorts.ProductAdminRepository.
func (p *productAdminRepositoryDB) FindCategoryByID(id uint) (*domain.Category, error) {
	var category domain.Category
	err := p.db.Where(&domain.Category{Model: gorm.Model{ID: id}}).First(&category)
	if err.Error != nil {
		return nil, err.Error
	}
	return &category, nil
}

// FindAllProducts implements adminPorts.ProductAdminRepository.
func (p *productAdminRepositoryDB) FindAllProducts() ([]domain.Product, error) {
	panic("unimplemented")
}

// FindProductByID implements adminPorts.ProductAdminRepository.
func (p *productAdminRepositoryDB) FindProductByID(id uint) (*domain.Product, error) {
	panic("unimplemented")
}

// UpdateProduct implements adminPorts.ProductAdminRepository.
func (p *productAdminRepositoryDB) UpdateProduct(product *domain.Product) error {
	panic("unimplemented")
}
