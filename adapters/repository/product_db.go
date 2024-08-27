package repository

import (
	"go-shop-api/common"
	"go-shop-api/core/domain"
	"go-shop-api/core/ports"
	"math"

	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) ports.ProductRepository {
	return &productRepositoryDB{db: db}
}

// FindAll implements ports.ProductRepository.
func (p *productRepositoryDB) FindAll(pagination *common.Pagination) (*common.Pagination, error) {
	var products []domain.Product

	// Get total count
	totalCount, err := common.GetTotalCount(p.db, &domain.Product{})
	if err != nil {
		return nil, err
	}

	// Apply pagination
	err = p.db.Preload("ProductImage").Scopes(common.Paginate(pagination)).Find(&products).Error
	if err != nil {
		return nil, err
	}

	// Set pagination values
	pagination.TotalRows = totalCount
	pagination.TotalPage = int(math.Ceil(float64(totalCount) / float64(pagination.GetLimit())))
	pagination.Items = products

	return pagination, nil
}
