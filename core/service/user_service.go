package service

import (
	"go-shop-api/adapters/errs"
	request "go-shop-api/core/model/resquest"
	"go-shop-api/core/ports"
)

type userService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{repo: repo}
}

// UpdateUserAvatar implements ports.UserService.
func (u *userService) UpdateUserAvatar(request *request.UpdateUserAvatarRequest) error {
	if request.ImageUrl == "" {
		return errs.NewBadRequestError("Image url is required")
	}

	user, err := u.repo.FindByUserId(request.UserId)
	if err != nil {
		return errs.NewNotFoundError("User not found")
	}

	user.Avatar = request.ImageUrl

	if err := u.repo.UpdateAvartar(user); err != nil {
		return errs.NewUnexpectedError("Failed to update user avatar")
	}

	return nil
}
