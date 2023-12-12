package user

import (
	"gorm.io/gorm"
)

type UserService interface {
}

type userService struct {
	// userRepository userRepo.UserRepository
}

func NewUserService(db *gorm.DB) userService {
	return userService{}
}
