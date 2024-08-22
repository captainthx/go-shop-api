package service

import (
	"errors"
	"go-shop-api/core/domain"
	"go-shop-api/core/ports"
	"go-shop-api/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo ports.AuthRepository
}

func NewAuthService(repo ports.AuthRepository) ports.AuthService {
	return &userService{repo: repo}
}

// CreateUser implements UserService.
func (u *userService) CreateUser(user *domain.User) error {
	if invalid, err := utils.InvalidName(user.Name); invalid || err != nil {
		return err
	}

	if invalid, err := utils.InvalidUsername(user.Username); invalid || err != nil {
		return err
	}

	if invalid, err := utils.InvalidPassword(user.Password); invalid || err != nil {
		return err
	}

	if invalid, err := utils.InvalidEmail(user.Email); invalid || err != nil {
		return err
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashPassword)

	if result := u.repo.Create(user); result != nil {
		return result
	}

	return nil
}

// LogIn implements ports.AuthService.
func (u *userService) LogIn(username string, password string) (*ports.LoginResponse, error) {
	if invalid, err := utils.InvalidUsername(username); invalid || err != nil {
		return nil, err
	}

	if invalid, err := utils.InvalidPassword(password); invalid || err != nil {
		return nil, err
	}

	resutl, err := u.repo.FindByUserName(username)

	if err != nil {
		return nil, err
	}

	// compare the password
	err = bcrypt.CompareHashAndPassword([]byte(resutl.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	loginResponse := &ports.LoginResponse{
		AccessToken:  generateAcessToken(resutl),
		RefreshToken: generateRefreshToken(resutl),
	}
	return loginResponse, nil
}

func generateRefreshToken(user *domain.User) string {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claims := refreshToken.Claims.(jwt.MapClaims)
	claims["issuer"] = os.Getenv("JWT_ISSUER")
	claims["auth"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return ""
	}
	return refreshTokenString
}

func generateAcessToken(user *domain.User) string {
	acessToken := jwt.New(jwt.SigningMethodHS256)
	claims := acessToken.Claims.(jwt.MapClaims)
	claims["issuer"] = os.Getenv("JWT_ISSUER")
	claims["auth"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	acessTokenString, err := acessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return ""
	}
	return acessTokenString
}
