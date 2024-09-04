package repository

import (
	"go-shop-api/core/domain"
	"go-shop-api/core/ports"

	"gorm.io/gorm"
)

type authRepositoryDB struct {
	db *gorm.DB
}

func NewAuthRepositoryDB(db *gorm.DB) ports.AuthRepository {
	return &authRepositoryDB{db: db}
}

func (r *authRepositoryDB) Create(user *domain.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// FindByUserName implements repository.UserRepository.
func (r *authRepositoryDB) FindByUserName(username string) (*domain.User, error) {
	var user domain.User
	result := r.db.Where(&domain.User{Username: username, Role: "customer"}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
