package dto

import (
	"time"

	"github.com/jakskal/koperasi-v2/internal/entity"
	"github.com/shopspring/decimal"
)

type CreateLoanInstallmentRequest struct {
	LoanID            int                        `json:"loan_id" binding:"required"`
	TransactionTypeID entity.LoanTransactionType `json:"transaction_type_id" binding:"required,oneof=1 2"`
	PaymentDate       time.Time                  `json:"payment_date" binding:"required"`
	PrincipalAmount   decimal.Decimal            `json:"principal_amount"`
	InterestAmount    decimal.Decimal            `json:"interest_amount"`
	Notes             string                     `json:"notes"`
}

func (s *CreateLoanInstallmentRequest) ToLoanInstallmentEntity() entity.LoanInstallment {
	return entity.LoanInstallment{
		LoanID:            s.LoanID,
		TransactionTypeID: s.TransactionTypeID,
		PaymentDate:       s.PaymentDate,
		PrincipalAmount:   s.PrincipalAmount,
		InterestAmount:    s.InterestAmount,
		Notes:             s.Notes,
	}
}

type UpdateLoanInstallmentRequest struct {
	ID                int
	TransactionTypeID entity.LoanTransactionType `json:"transaction_type_id" binding:"required,oneof=1 2"`
	PaymentDate       time.Time                  `json:"payment_date" binding:"required"`
	PrincipalAmount   decimal.Decimal            `json:"principal_amount"`
	InterestAmount    decimal.Decimal            `json:"interest_amount"`
	Notes             string                     `json:"notes"`
	ChangeNotes       string                     `json:"change_notes"`
}

type GetLoanInstallmentListRequest struct {
	LoanID   *int   `form:"loan_id"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	OrderBy  string `form:"order_by"`
	Order    string `form:"order"`
}

type GetQueryLoanInstallmentListResponse struct {
	LoanInstallments []entity.LoanInstallment `json:"embedded"`
	Pagination       BasePaginationResult     `json:"pagination"`
}

type GetLoanInstallmentListResponse struct {
	LoanInstallments []GetLoanInstallmentResponse `json:"embedded"`
	Pagination       BasePaginationResult         `json:"pagination"`
}

type GetLoanInstallmentResponse struct {
	ID                     int                        `json:"id"`
	LoanID                 int                        `json:"loan_id"`
	TransactionTypeID      entity.LoanTransactionType `json:"transaction_type_id"`
	PaymentDate            time.Time                  `json:"payment_date"`
	PrincipalAmount        float64                    `json:"principal_amount"`
	InterestAmount         float64                    `json:"interest_amount"`
	Notes                  string                     `json:"notes"`
	entity.TimeDefault     `gorm:"embedded"`
	LoanInstallmentChanges []entity.LoanInstallmentChange `json:"installment_changes"`
}

func (i *GetLoanInstallmentResponse) FromEntity(in entity.LoanInstallment) {
	i.ID = in.ID
	i.LoanID = in.LoanID
	i.TransactionTypeID = in.TransactionTypeID
	i.PaymentDate = in.PaymentDate
	i.PrincipalAmount = in.PrincipalAmount.InexactFloat64()
	i.InterestAmount = in.InterestAmount.InexactFloat64()
	i.Notes = in.Notes
	i.TimeDefault = in.TimeDefault
	i.LoanInstallmentChanges = in.LoanInstallmentChanges
}
