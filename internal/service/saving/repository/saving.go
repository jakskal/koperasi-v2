package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jakskal/koperasi-v2/internal/entity"
	"github.com/jakskal/koperasi-v2/pkg/dto"
	"github.com/jakskal/koperasi-v2/pkg/middleware"
	"github.com/jakskal/koperasi-v2/pkg/paginator"
	"github.com/jakskal/koperasi-v2/pkg/token"
	"gorm.io/gorm"
)

type savingRepository struct {
	db *gorm.DB
}

func NewSavingRepository(db *gorm.DB) SavingRepository {
	return &savingRepository{
		db: db,
	}
}

func (r *savingRepository) Create(ctx context.Context, req entity.Saving) error {
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

func (r *savingRepository) Get(ctx context.Context, ID int) (entity.Saving, error) {
	var saving entity.Saving
	if err := r.db.Where(&entity.Saving{ID: ID}).Where("deleted_at is NULL").First(&saving).Error; err != nil {
		return saving, err
	}
	return saving, nil
}

func (r *savingRepository) GetUserSavingDetail(ctx context.Context, ID int, userID int) (entity.Saving, error) {
	var saving entity.Saving
	if err := r.db.Where(&entity.Saving{ID: ID, UserID: userID}).Where("deleted_at is NULL").First(&saving).Error; err != nil {
		return saving, err
	}
	return saving, nil
}

func (r *savingRepository) List(ctx context.Context, req dto.GetSavingListRequest) (dto.GetSavingListResponse, error) {
	var (
		result  dto.GetSavingListResponse
		savings []entity.Saving
		count   int64
	)
	q := r.db.Model(&entity.Saving{}).Preload("SavingChanges").Where("savings.deleted_at is NULL")

	if req.TypeID != nil {
		q = q.Where("saving_type_id = ?", req.TypeID)
	}

	if req.UserID != nil {
		q = q.Where("user_id = ?", req.UserID)
	}

	q = q.Scopes(paginator.PaginateGin(req.Page, req.PageSize))
	q.Count(&count)
	resultQuery := q.Find(&savings)
	if err := resultQuery.Error; err != nil {
		return result, err
	}

	result.Savings = savings
	result.Pagination = dto.BasePaginationResult{
		Page:     req.Page,
		PageSize: req.PageSize,
		Count:    int(count),
	}

	fmt.Println(result)
	return result, nil
}

func (r *savingRepository) Update(ctx context.Context, req dto.UpdateSavingRequest) error {
	var currentSaving entity.Saving

	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := r.db.Model(entity.Saving{}).Where("id = ? and deleted_at is null", req.ID).First(&currentSaving).Error
	if err != nil {
		return err
	}
	tokenInfo := ctx.Value(middleware.TokenInfoContextKey).(token.Claims)

	savingChanges := entity.SavingChange{
		SavingID:          currentSaving.ID,
		TransactionTypeID: currentSaving.TransactionTypeID,
		Notes:             currentSaving.Notes,
		Amount:            currentSaving.Amount,
		ChangesNotes:      req.ChangeNotes,
		CreatedAt:         currentSaving.CreatedAt,
		CreatedBy:         tokenInfo.UserID,
	}

	if err := tx.Omit("SavingChanges").Updates(entity.Saving{
		ID:                req.ID,
		Amount:            req.Amount,
		Notes:             req.Notes,
		TransactionTypeID: req.TransactionTypeID,
		TimeDefault:       entity.TimeDefault{UpdatedBy: &tokenInfo.UserID},
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&savingChanges).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (r *savingRepository) Delete(ctx context.Context, ID int) error {

	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	tokenInfo := ctx.Value(middleware.TokenInfoContextKey).(token.Claims)

	timeNow := time.Now()
	if err := tx.Omit("SavingChanges").Updates(entity.Saving{
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
