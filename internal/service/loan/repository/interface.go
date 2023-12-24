package repository

import (
	"context"

	"github.com/jakskal/koperasi-v2/internal/entity"
	"github.com/jakskal/koperasi-v2/pkg/dto"
)

type LoanRepository interface {
	Create(context.Context, entity.Loan) error
	Get(ctx context.Context, ID int) (entity.Loan, error)
	GetUserLoanDetail(ctx context.Context, ID int, userID int) (entity.Loan, error)
	List(ctx context.Context, req dto.GetLoanListRequest) (dto.GetQueryLoanListResponse, error)
	Update(context.Context, dto.UpdateLoanRequest) error
	Delete(ctx context.Context, ID int) error
	CreateLoanType(ctx context.Context, req entity.LoanType) error
	UpdateLoanType(ctx context.Context, req dto.UpdateLoanTypeRequest) error
	ListLoanType(ctx context.Context) ([]entity.LoanType, error)
	GetLoanType(ctx context.Context, ID int) (entity.LoanType, error)
	CreateLoanInstallment(context.Context, entity.LoanInstallment) error
	ListLoanInstallment(context.Context, dto.GetLoanInstallmentListRequest) (dto.GetQueryLoanInstallmentListResponse, error)
	GetLoanInstallment(ctx context.Context, ID int) (entity.LoanInstallment, error)
	UpdateLoanInstallment(context.Context, dto.UpdateLoanInstallmentRequest) error
	DeleteLoanInstallment(ctx context.Context, ID int) error
}
