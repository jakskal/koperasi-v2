package repository

import (
	"context"

	"github.com/jakskal/koperasi-v2/internal/entity"
	"github.com/jakskal/koperasi-v2/pkg/dto"
	"github.com/jakskal/koperasi-v2/pkg/middleware"
	"github.com/jakskal/koperasi-v2/pkg/token"
)

func (r *loanRepository) CreateLoanType(ctx context.Context, req entity.LoanType) error {
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

func (r *loanRepository) UpdateLoanType(ctx context.Context, req dto.UpdateLoanTypeRequest) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	tokenInfo := ctx.Value(middleware.TokenInfoContextKey).(token.Claims)

	if err := tx.Updates(&entity.LoanType{
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
func (r *loanRepository) ListLoanType(ctx context.Context) ([]entity.LoanType, error) {
	var result []entity.LoanType
	if err := r.db.Model(entity.LoanType{}).Where("deleted_at is null").Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

func (r *loanRepository) GetLoanType(ctx context.Context, ID int) (entity.LoanType, error) {
	var result entity.LoanType
	if err := r.db.Where(&entity.LoanType{ID: ID}).Where("deleted_at is NULL").First(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
