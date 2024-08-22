package ports

import "go-shop-api/core/domain"

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}

type AuthService interface {
	CreateUser(user *domain.User) error
	LogIn(username string, password string) (*LoginResponse, error)
}

type AuthRepository interface {
	Create(user *domain.User) error
	FindByUserName(username string) (*domain.User, error)
}

type UserRepository interface {
	FindByID(id uint) (*domain.User, error)
	FindOrderByUser(user *domain.User) ([]domain.Order, error)
}
