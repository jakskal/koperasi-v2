package repository

import (
	"context"

	"github.com/jakskal/koperasi-v2/internal/entity"
	"github.com/jakskal/koperasi-v2/pkg/dto"
	"github.com/jakskal/koperasi-v2/pkg/middleware"
	"github.com/jakskal/koperasi-v2/pkg/token"
)

func (r *savingRepository) CreateSavingTypes(ctx context.Context, req entity.SavingType) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&req).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *savingRepository) UpdateSavingType(ctx context.Context, req dto.UpdateSavingTypeRequest) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	tokenInfo := ctx.Value(middleware.TokenInfoContextKey).(token.Claims)

	if err := tx.Updates(&entity.SavingType{
		ID: req.ID, Name: req.Name,
		TimeDefault: entity.TimeDefault{
			UpdatedBy: &tokenInfo.UserID,
		},
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *savingRepository) ListSavingType(ctx context.Context) ([]entity.SavingType, error) {
	var result []entity.SavingType

	if err := r.db.Model(entity.SavingType{}).Where("deleted_at is null").Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
