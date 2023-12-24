package dto

import "github.com/jakskal/koperasi-v2/internal/entity"

type CreateSavingRequest struct {
	UserID            int                           `json:"user_id" binding:"required"`
	SavingTypeID      int                           `json:"saving_type_id" binding:"required"`
	TransactionTypeID *entity.SavingTransactionType `json:"transaction_type_id" binding:"required,oneof=1 2"`
	Amount            int                           `json:"amount"`
	Notes             string                        `json:"notes"`
}

func (s *CreateSavingRequest) ToSavingEntity() entity.Saving {
	return entity.Saving{
		UserID:            s.UserID,
		SavingTypeID:      s.SavingTypeID,
		TransactionTypeID: *s.TransactionTypeID,
		Amount:            s.Amount,
		Notes:             s.Notes,
	}
}

type UpdateSavingRequest struct {
	ID                int
	Amount            int                          `json:"amount"`
	Notes             string                       `json:"notes"`
	ChangeNotes       string                       `json:"change_notes"`
	TransactionTypeID entity.SavingTransactionType `json:"transaction_type_id" binding:"oneof=1 2"`
}

type GetSavingListRequest struct {
	TypeID   *int   `form:"type_id"`
	UserID   *int   `form:"user_id"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	OrderBy  string `form:"order_by"`
	Order    string `form:"order"`
}

type GetSavingListResponse struct {
	Savings    []entity.Saving      `json:"embedded"`
	Pagination BasePaginationResult `json:"pagination"`
}
