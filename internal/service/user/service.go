package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/jakskal/koperasi-v2/internal/entity"
	userRepo "github.com/jakskal/koperasi-v2/internal/service/user/repository"
	"github.com/jakskal/koperasi-v2/pkg/dto"
	"github.com/jakskal/koperasi-v2/pkg/hash"
	"github.com/jakskal/koperasi-v2/pkg/token"
	"gorm.io/gorm"
)

type UserService interface {
	Register(context.Context, dto.RegisterRequest) error
	Login(context.Context, dto.LoginRequest) (dto.LoginResponse, error)
	UpsertAttribute(context.Context, entity.UserAttribute) (entity.UserAttribute, error)
	CreateUser(context.Context, dto.CreateUserRequest) (entity.User, error)
	GetUser(ctx context.Context, id int) (entity.User, error)
	GetUsers(ctx context.Context, req dto.GetUsersRequest) (dto.GetUserResponse, error)
	UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userService struct {
	userRepo userRepo.UserRepository
	tokenSvc token.Service
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		userRepo: userRepo.NewUserRepository(db),
		tokenSvc: *token.NewService(),
	}
}

func (s *userService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (entity.User, error) {
	var user entity.User
	hashedPassword, err := hash.Password(req.Password)
	if err != nil {
		return user, nil
	}

	req.Password = hashedPassword
	req.StatusID = entity.USER_STATUS_VERIVIED
	user = req.ToUserEntity()
	err = s.userRepo.SaveUser(ctx, &user)
	if err != nil {
		return user, errors.New(fmt.Sprintf("failed create user, err: %s", err.Error()))
	}
	return user, nil
}

func (s *userService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	var res dto.LoginResponse

	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return res, err
	}

	v := hash.ComparePasswords(user.Password, req.Password)
	if v == false {
		return res, errors.New("invalid password")
	}

	token, err := s.tokenSvc.CreateToken(ctx, &token.CreateTokenRequest{
		UserID: user.ID,
		RoleID: int(user.RoleID),
	})
	if err != nil {
		return res, err
	}

	res.Token = token.AccessToken
	return res, nil
}

func (s *userService) UpsertAttribute(context.Context, entity.UserAttribute) (entity.UserAttribute, error) {
	return entity.UserAttribute{}, nil
}

func (s *userService) Register(context.Context, dto.RegisterRequest) error {
	return nil
}

func (s *userService) GetUser(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	user, err := s.userRepo.GetById(ctx, id)
	if err != nil {
		return user, dto.ErrorWrap("failed get user by id", err)
	}

	return user, nil

}

func (s *userService) GetUsers(ctx context.Context, req dto.GetUsersRequest) (dto.GetUserResponse, error) {
	var user dto.GetUserResponse
	user, err := s.userRepo.GetUsers(ctx, req)
	if err != nil {
		return user, dto.ErrorWrap("failed get user by id", err)
	}

	return user, nil

}

func (s *userService) UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (entity.User, error) {
	var user entity.User

	user = req.ToUserEntity()
	err := s.userRepo.SaveUser(ctx, &user)
	if err != nil {
		return user, errors.New(fmt.Sprintf("failed update user, err: %s", err.Error()))
	}
	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int) error {
	err := s.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return errors.New(fmt.Sprintf("failed delete user, err: %s", err.Error()))
	}
	return nil
}
