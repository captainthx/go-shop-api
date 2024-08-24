package adminService

import (
	"go-shop-api/adapters/errs"
	"go-shop-api/core/domain"
	"go-shop-api/core/ports"
	adminPorts "go-shop-api/core/ports/admin"
	"go-shop-api/logs"
	"go-shop-api/utils"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authAdminService struct {
	repo adminPorts.AuthAdminRepository
}

func NewAuthAdminService(repo adminPorts.AuthAdminRepository) adminPorts.AuthAdminService {
	return &authAdminService{repo: repo}
}

// CreateAdmin implements adminPorts.AuthAdminService.
func (a *authAdminService) CreateAdmin(user *domain.User) error {
	if invalid, err := utils.InvalidName(user.Name); invalid || err != nil {
		logs.Error(err)
		return errs.AppError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	if invalid, err := utils.InvalidUsername(user.Username); invalid || err != nil {
		logs.Error(err)
		return errs.NewBadRequestError(err.Error())
	}

	if invalid, err := utils.InvalidPassword(user.Password); invalid || err != nil {
		logs.Error(err)
		return errs.NewBadRequestError(err.Error())
	}

	if invalid, err := utils.InvalidEmail(user.Email); invalid || err != nil {
		logs.Error(err)
		return errs.NewBadRequestError(err.Error())
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		logs.Error(err)
		return errs.NewBadRequestError("Invalid password")
	}
	user.Password = string(hashPassword)
	user.Role = domain.Role("admin")

	err = a.repo.CreateAdmin(user)
	if err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError("unexpected error")
	}

	return nil
}

// LogIn implements adminPorts.AuthAdminService.
func (a *authAdminService) LogIn(username string, password string) (*ports.LoginResponse, error) {
	if invalid, err := utils.InvalidUsername(username); invalid || err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	if invalid, err := utils.InvalidPassword(password); invalid || err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	resutl, err := a.repo.FindByUserName(username)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("Admin not found")
		}
		logs.Error(err)

		return nil, errs.NewUnexpectedError("unexpected error")
	}

	// compare the password
	err = bcrypt.CompareHashAndPassword([]byte(resutl.Password), []byte(password))
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError("Invalid password")
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
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()

	acessTokenString, err := acessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return ""
	}
	return acessTokenString
}
