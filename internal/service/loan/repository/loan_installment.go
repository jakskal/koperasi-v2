package repository

import (
	"context"
	"time"

	"github.com/jakskal/koperasi-v2/internal/entity"
	"github.com/jakskal/koperasi-v2/pkg/dto"
	"github.com/jakskal/koperasi-v2/pkg/middleware"
	"github.com/jakskal/koperasi-v2/pkg/paginator"
	"github.com/jakskal/koperasi-v2/pkg/token"
)

func (r *loanRepository) CreateLoanInstallment(ctx context.Context, req entity.LoanInstallment) error {
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

	tx.Commit()
	return nil
}
func (r *loanRepository) ListLoanInstallment(ctx context.Context, req dto.GetLoanInstallmentListRequest) (dto.GetQueryLoanInstallmentListResponse, error) {
	var (
		result           dto.GetQueryLoanInstallmentListResponse
		loanInstallments []entity.LoanInstallment
		count            int64
	)
	q := r.db.Model(&entity.LoanInstallment{}).Preload("LoanInstallmentChanges").Where("loan_installments.deleted_at is NULL")

	if req.LoanID != nil {
		q = q.Where("loan_id = ?", req.LoanID)
	}

	q = q.Scopes(paginator.PaginateGin(req.Page, req.PageSize))
	q.Count(&count)
	resultQuery := q.Find(&loanInstallments)
	if err := resultQuery.Error; err != nil {
		return result, err
	}

	result.LoanInstallments = loanInstallments
	result.Pagination = dto.BasePaginationResult{
		Page:     req.Page,
		PageSize: req.PageSize,
		Count:    int(count),
	}

	return result, nil
}
func (r *loanRepository) GetLoanInstallment(ctx context.Context, ID int) (entity.LoanInstallment, error) {
	var result entity.LoanInstallment
	if err := r.db.Where(&entity.LoanInstallment{ID: ID}).Where("deleted_at is NULL").First(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
func (r *loanRepository) UpdateLoanInstallment(ctx context.Context, req dto.UpdateLoanInstallmentRequest) error {
	var currentInstallment entity.LoanInstallment

	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := r.db.Model(entity.LoanInstallment{}).Where("id = ? and deleted_at is null", req.ID).First(&currentInstallment).Error
	if err != nil {
		return err
	}
	tokenInfo := ctx.Value(middleware.TokenInfoContextKey).(token.Claims)
	installmentChanges := entity.LoanInstallmentChange{
		LoanInstallmentID: currentInstallment.ID,
		TransactionTypeID: currentInstallment.TransactionTypeID,
		PaymentDate:       currentInstallment.PaymentDate,
		PrincipalAmount:   currentInstallment.PrincipalAmount,
		InterestAmount:    currentInstallment.InterestAmount,
		Notes:             currentInstallment.Notes,
		ChangesNotes:      req.ChangeNotes,
		CreatedBy:         tokenInfo.UserID,
	}

	if err := tx.Omit("LoanChanges").Updates(entity.LoanInstallment{
		ID:                req.ID,
		TransactionTypeID: req.TransactionTypeID,
		PaymentDate:       req.PaymentDate,
		PrincipalAmount:   req.PrincipalAmount,
		InterestAmount:    req.InterestAmount,
		Notes:             req.Notes,
		TimeDefault:       entity.TimeDefault{UpdatedBy: &tokenInfo.UserID},
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&installmentChanges).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
func (r *loanRepository) DeleteLoanInstallment(ctx context.Context, ID int) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	tokenInfo := ctx.Value(middleware.TokenInfoContextKey).(token.Claims)

	timeNow := time.Now()
	if err := tx.Omit("LoanInstallmentChanges").Updates(entity.LoanInstallment{
		ID: ID,
		TimeDefault: entity.TimeDefault{
			UpdatedBy: &tokenInfo.UserID,
			DeletedAt: &timeNow,
			DeletedBy: &tokenInfo.UserID},
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
