package dto

import (
	"github.com/jakskal/koperasi-v2/internal/entity"
	"github.com/shopspring/decimal"
)

type CreateLoanRequest struct {
	UserID               int             `json:"user_id" binding:"required"`
	Name                 string          `json:"name" binding:"required"`
	Amount               decimal.Decimal `json:"amount" binding:"required"`
	LoanTypeID           int             `json:"loan_type_id" binding:"required"`
	RatioPercentage      decimal.Decimal
	TotalRatioAmount     decimal.Decimal
	InstallmentQtyTarget int    `json:"installment_qty_target" binding:"required"`
	Notes                string `json:"notes"`
}

func (s *CreateLoanRequest) WithInterest(in entity.LoanType) *CreateLoanRequest {
	percentage := decimal.NewFromFloat(0.01)
	s.RatioPercentage = in.RatioPercentage
	s.TotalRatioAmount = s.Amount.Mul(in.RatioPercentage.Mul(percentage))
	return s
}

func (s *CreateLoanRequest) ToLoanEntity() entity.Loan {
	return entity.Loan{
		UserID:               s.UserID,
		LoanTypeID:           s.LoanTypeID,
		Name:                 s.Name,
		Amount:               s.Amount,
		RatioPercentage:      s.RatioPercentage,
		TotalRatioAmount:     s.TotalRatioAmount,
		InstallmentQtyTarget: s.InstallmentQtyTarget,
		Notes:                s.Notes,
	}
}

type UpdateLoanRequest struct {
	ID                   int
	Name                 string          `json:"name"`
	Amount               decimal.Decimal `json:"amount" binding:"required"`
	RatioPercentage      decimal.Decimal `json:"ratio_percentage"`
	TotalRatioAmount     decimal.Decimal `json:"total_ratio_amount"`
	InstallmentQtyTarget int             `json:"installment_qty_target"`
	Notes                string          `json:"notes"`
	ChangeNotes          string          `json:"change_notes"`
}

type GetLoanListRequest struct {
	TypeID   *int   `form:"type_id"`
	UserID   *int   `form:"user_id"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	OrderBy  string `form:"order_by"`
	Order    string `form:"order"`
}

type GetQueryLoanListResponse struct {
	Loans      []entity.Loan        `json:"embedded"`
	Pagination BasePaginationResult `json:"pagination"`
}

type GetLoanListResponse struct {
	Loans      []GetLoanResponse    `json:"embedded"`
	Pagination BasePaginationResult `json:"pagination"`
}

type GetLoanResponse struct {
	ID                   int     `json:"id"`
	UserID               int     `json:"user_id"`
	LoanTypeID           int     `json:"loan_type_id"`
	Name                 string  `json:"name"`
	RatioPercentage      float64 `json:"ratio_percentage"`
	Amount               int64   `json:"amount"`
	TotalRatioAmount     int64   `json:"total_ratio_amount"`
	InstallmentQtyTarget int     `json:"installment_qty_target"`
	Notes                string  `json:"notes"`
	entity.TimeDefault   `gorm:"embedded"`
	LoanChanges          []entity.LoanChange `gorm:"foreignKey:LoanID;references:ID"`
}

func (l *GetLoanResponse) FromEntity(in entity.Loan) {
	l.ID = in.ID
	l.UserID = in.UserID
	l.LoanTypeID = in.LoanTypeID
	l.Name = in.Name
	l.Amount = in.Amount.IntPart()
	l.RatioPercentage = in.RatioPercentage.InexactFloat64()
	l.TotalRatioAmount = in.TotalRatioAmount.IntPart()
	l.InstallmentQtyTarget = in.InstallmentQtyTarget
	l.Notes = in.Notes
	l.TimeDefault = in.TimeDefault
	l.LoanChanges = in.LoanChanges
}
