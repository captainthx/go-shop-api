package repository

import (
	"go-shop-api/core/domain"
	"go-shop-api/core/ports"

	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) ports.UserRepository {
	return &userRepositoryDB{db: db}
}

// FindByUserId implements ports.UserRepository.
func (u *userRepositoryDB) FindByUserId(id uint) (*domain.User, error) {
	var user domain.User
	err := u.db.Model(&domain.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindOrderByUser implements ports.UserRepository.
func (u *userRepositoryDB) FindOrderByUser(user *domain.User) ([]domain.Order, error) {
	var orders []domain.Order

	err := u.db.Model(&domain.Order{}).Where("user_id = ?", user.ID).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// UpdateAvartar implements ports.UserRepository.
func (u *userRepositoryDB) UpdateAvartar(user *domain.User) error {
	err := u.db.Model(&domain.User{}).Where("id = ?", user.ID).Update("avatar", user.Avatar).Error
	if err != nil {
		return err
	}
	return nil
}
