package repository

import (
	"context"

	"github.com/jakskal/koperasi-v2/internal/entity"
	"github.com/jakskal/koperasi-v2/pkg/dto"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *entity.User) error
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	GetById(ctx context.Context, id int) (entity.User, error)
	GetUsers(ctx context.Context, req dto.GetUsersRequest) (dto.GetUserResponse, error)
	DeleteUser(ctx context.Context, id int) error
}
