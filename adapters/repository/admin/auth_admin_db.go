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
	var user domain.User
	result := a.db.Where(&domain.User{Username: username, Role: "admin"}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
