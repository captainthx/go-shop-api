package adminService

import (
	"go-shop-api/adapters/errs"
	"go-shop-api/core/domain"
	adminPorts "go-shop-api/core/ports/admin"
	"go-shop-api/logs"
	"go-shop-api/utils"

	"gorm.io/gorm"
)

type productAdminService struct {
	repo adminPorts.ProductAdminRepository
}

func NewProductAdminService(repo adminPorts.ProductAdminRepository) adminPorts.ProductAdminService {
	return &productAdminService{repo: repo}
}

// CreateProduct implements adminPorts.ProductAdminService.
func (p *productAdminService) CreateProduct(product *domain.Product) error {

	if invalid, err := utils.InvalidProductName(product.Name); invalid || err != nil {
		logs.Error(err)
		return errs.NewBadRequestError(err.Error())
	}

	if invalid, err := utils.InvalidProductPrice(product.Price); invalid || err != nil {
		logs.Error(err)
		return errs.NewBadRequestError(err.Error())
	}

	if invalid, err := utils.InvalidQuantity(product.Quantity); invalid || err != nil {
		logs.Error(err)
		return errs.NewBadRequestError(err.Error())
	}

	var productImage []string
	for _, img := range product.ProductImage {
		productImage = append(productImage, img.URL)
	}

	if invalid, err := utils.InvalidProductImage(productImage); invalid || err != nil {
		logs.Error(err)
		return errs.NewBadRequestError(err.Error())
	}

	category, err := p.repo.FindCategoryByID(product.CategoryID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logs.Error(err)
			return errs.NewNotFoundError("CategoryID not found")
		}
		logs.Error(err)
		return errs.NewUnexpectedError(err.Error())
	}

	product.CategoryID = category.ID

	if err := p.repo.CreateProduct(product); err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError(err.Error())
	}

	return nil
}
