package adminService

import (
	"go-shop-api/adapters/errs"
	"go-shop-api/core/domain"
	adminPorts "go-shop-api/core/ports/admin"
	"go-shop-api/logs"
	"go-shop-api/utils"
)

type categoryAdminService struct {
	repo adminPorts.CategoryAdminRepository
}

func NewCategoryAdminService(repo adminPorts.CategoryAdminRepository) adminPorts.CategoryAdminService {
	return &categoryAdminService{repo: repo}
}

// CreateCategory implements adminPorts.CategoryAdminService.
func (c *categoryAdminService) CreateCategory(category *domain.Category) error {

	if invalid, err := utils.InvalidCategoryName(category.Name); invalid || err != nil {
		logs.Error(err)
		return errs.NewBadRequestError(err.Error())
	}
	err := c.repo.CreateCategory(category)

	if err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}

// GetCategory implements adminPorts.CategoryAdminService.
func (c *categoryAdminService) GetCategory() ([]domain.Category, error) {
	result, err := c.repo.FindAllCategory()

	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return result, nil
}
