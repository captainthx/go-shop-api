package adminRepository

import (
	"go-shop-api/core/domain"
	adminPorts "go-shop-api/core/ports/admin"

	"gorm.io/gorm"
)

type authAdminRepositoryDB struct {
	db *gorm.DB
}

func NewAuthAdminRepositoryDB(db *gorm.DB) adminPorts.AuthAdminRepository {
	return &authAdminRepositoryDB{db: db}
}

// CreateAdmin implements adminPorts.AuthAdminRepository.
func (a *authAdminRepositoryDB) CreateAdmin(user *domain.User) error {
	if err := a.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// FindByUserName implements adminPorts.AuthAdminRepository.
func (a *authAdminRepositoryDB) FindByUserName(username string) (*domain.User, error) {
	result := a.db.Where(&domain.User{Username: username}).First(&domain.User{})
	if result.Error != nil {
		return nil, result.Error
	}
	var user domain.User
	if err := result.Scan(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
