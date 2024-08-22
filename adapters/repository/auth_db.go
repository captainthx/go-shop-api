package repository

import (
	"errors"
	"go-shop-api/core/domain"
	"go-shop-api/core/ports"

	"gorm.io/gorm"
)

type authRepositoryDB struct {
	db *gorm.DB
}

func NewauthRepositoryDB(db *gorm.DB) ports.AuthRepository {
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
	result := r.db.Where(&domain.User{Username: username}).First(&domain.User{})
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	var user domain.User
	if err := result.Scan(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
