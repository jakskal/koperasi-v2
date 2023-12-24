package repository

import (
	"context"
	"time"

	"github.com/jakskal/koperasi-v2/internal/entity"
	"github.com/jakskal/koperasi-v2/pkg/dto"
	"github.com/jakskal/koperasi-v2/pkg/middleware"
	"github.com/jakskal/koperasi-v2/pkg/paginator"
	"github.com/jakskal/koperasi-v2/pkg/token"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) SaveUser(ctx context.Context, req *entity.User) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if req.ID == 0 {
		if err := tx.Create(req).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if err := tx.Omit("UserAttribute").Updates(req).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.Updates(req.UserAttribute).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where(&entity.User{Email: email}).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetById(ctx context.Context, id int) (entity.User, error) {

	var user entity.User
	if err := r.db.Where(&entity.User{ID: id}).Where("deleted_at is NULL").First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUsers(ctx context.Context, req dto.GetUsersRequest) (dto.GetUserResponse, error) {
	var (
		result dto.GetUserResponse
		users  []entity.User
		count  int64
	)
	q := r.db.Model(&entity.User{}).Preload("UserAttribute").Where("users.deleted_at is NULL")

	if req.RoleID != nil {
		q = q.Where("role_id = ?", req.RoleID)
	}

	q = q.Scopes(paginator.PaginateGin(req.Page, req.PageSize))
	q.Count(&count)
	if err := q.Find(&users).Error; err != nil {
		return result, err
	}

	result.Users = users
	result.Pagination = dto.BasePaginationResult{
		Page:     req.Page,
		PageSize: req.PageSize,
		Count:    int(count),
	}
	return result, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, ID int) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	tokenInfo := ctx.Value(middleware.TokenInfoContextKey).(token.Claims)
	t := time.Now()
	if err := tx.Updates(&entity.User{ID: ID, TimeDefault: entity.TimeDefault{
		DeletedAt: &t,
		UpdatedBy: &tokenInfo.UserID,
	}}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
