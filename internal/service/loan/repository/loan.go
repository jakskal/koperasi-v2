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

type loanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(DB *gorm.DB) LoanRepository {
	return &loanRepository{
		db: DB,
	}
}

func (r *loanRepository) Create(ctx context.Context, req entity.Loan) error {
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
func (r *loanRepository) Get(ctx context.Context, ID int) (entity.Loan, error) {
	var result entity.Loan
	if err := r.db.Where(&entity.Loan{ID: ID}).Where("deleted_at is NULL").First(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
func (r *loanRepository) GetUserLoanDetail(ctx context.Context, ID int, userID int) (entity.Loan, error) {
	var result entity.Loan
	if err := r.db.Where(&entity.Loan{ID: ID, UserID: userID}).Where("deleted_at is NULL").First(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
func (r *loanRepository) List(ctx context.Context, req dto.GetLoanListRequest) (dto.GetQueryLoanListResponse, error) {
	var (
		result dto.GetQueryLoanListResponse
		loans  []entity.Loan
		count  int64
	)
	q := r.db.Model(&entity.Loan{}).Preload("LoanChanges").Where("loans.deleted_at is NULL")

	if req.TypeID != nil {
		q = q.Where("loan_type_id = ?", req.TypeID)
	}

	if req.UserID != nil {
		q = q.Where("user_id = ?", req.UserID)
	}

	q = q.Scopes(paginator.PaginateGin(req.Page, req.PageSize))
	q.Count(&count)
	resultQuery := q.Find(&loans)
	if err := resultQuery.Error; err != nil {
		return result, err
	}

	result.Loans = loans
	result.Pagination = dto.BasePaginationResult{
		Page:     req.Page,
		PageSize: req.PageSize,
		Count:    int(count),
	}

	return result, nil
}
func (r *loanRepository) Update(ctx context.Context, req dto.UpdateLoanRequest) error {
	var currentLoan entity.Loan

	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := r.db.Model(entity.Loan{}).Where("id = ? and deleted_at is null", req.ID).First(&currentLoan).Error
	if err != nil {
		return err
	}
	tokenInfo := ctx.Value(middleware.TokenInfoContextKey).(token.Claims)

	loanChanges := entity.LoanChange{
		LoanID:           currentLoan.ID,
		Name:             currentLoan.Name,
		Amount:           currentLoan.Amount,
		RatioPercentage:  currentLoan.RatioPercentage,
		TotalRatioAmount: currentLoan.TotalRatioAmount,
		Notes:            currentLoan.Notes,
		ChangesNotes:     req.ChangeNotes,
		CreatedBy:        tokenInfo.UserID,
	}

	if err := tx.Omit("LoanChanges").Updates(entity.Loan{
		ID:                   req.ID,
		Name:                 req.Name,
		Amount:               req.Amount,
		RatioPercentage:      req.RatioPercentage,
		TotalRatioAmount:     req.TotalRatioAmount,
		InstallmentQtyTarget: req.InstallmentQtyTarget,
		Notes:                req.Notes,
		TimeDefault:          entity.TimeDefault{UpdatedBy: &tokenInfo.UserID},
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&loanChanges).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
func (r *loanRepository) Delete(ctx context.Context, ID int) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	tokenInfo := ctx.Value(middleware.TokenInfoContextKey).(token.Claims)

	timeNow := time.Now()
	if err := tx.Omit("SavingChanges").Updates(entity.Loan{
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
