package adminPorts

import (
	"go-shop-api/core/domain"
	"go-shop-api/core/ports"
)

type AuthAdminService interface {
	CreateAdmin(user *domain.User) error
	LogIn(username string, password string) (*ports.LoginResponse, error)
}

type AuthAdminRepository interface {
	CreateAdmin(user *domain.User) error
	FindByUserName(username string) (*domain.User, error)
}
