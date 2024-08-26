package adminPorts

import (
	"go-shop-api/core/domain"
	"go-shop-api/core/model/response"
)

type AuthAdminService interface {
	CreateAdmin(user *domain.User) error
	LogIn(username string, password string) (*response.LoginResponse, error)
}

type AuthAdminRepository interface {
	CreateAdmin(user *domain.User) error
	FindByUserName(username string) (*domain.User, error)
}

type ProductAdminRepository interface {
	CreateProduct(product *domain.Product) error
	FindProductByID(id uint) (*domain.Product, error)
	FindCategoryByID(id uint) (*domain.Category, error)
	FindAllProducts() ([]domain.Product, error)
	UpdateProduct(product *domain.Product) error
	DeleteProduct(product *domain.Product) error
}

type ProductAdminService interface {
	CreateProduct(product *domain.Product) error
}

type CategoryAdminRepository interface {
	CreateCategory(category *domain.Category) error
	FindAllCategory() ([]domain.Category, error)
}

type CategoryAdminService interface {
	CreateCategory(category *domain.Category) error
	GetCategory() ([]response.CategoryResponse, error)
}
