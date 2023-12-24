package loan

import (
	"context"

	"github.com/jakskal/koperasi-v2/internal/entity"
	"github.com/jakskal/koperasi-v2/internal/service/loan/repository"
	"github.com/jakskal/koperasi-v2/pkg/dto"
	"github.com/jakskal/koperasi-v2/pkg/middleware"
	"github.com/jakskal/koperasi-v2/pkg/token"
	"gorm.io/gorm"
)

type LoanService interface {
	Create(context.Context, dto.CreateLoanRequest) error
	List(context.Context, dto.GetLoanListRequest) (dto.GetLoanListResponse, error)
	Get(ctx context.Context, ID int) (entity.Loan, error)
	Update(context.Context, dto.UpdateLoanRequest) error
	Delete(ctx context.Context, ID int) error
	CreateLoanType(ctx context.Context, req dto.CreateLoanTypeRequest) error
	UpdateLoanType(context.Context, dto.UpdateLoanTypeRequest) error
	ListLoanType(context.Context) ([]dto.GetLoanTypeResponse, error)
	CreateLoanInstallment(context.Context, dto.CreateLoanInstallmentRequest) error
	GetLoanInstallment(ctx context.Context, ID int) (entity.LoanInstallment, error)
	ListLoanInstallment(context.Context, dto.GetLoanInstallmentListRequest) (dto.GetLoanInstallmentListResponse, error)
	UpdateLoanInstallment(context.Context, dto.UpdateLoanInstallmentRequest) error
	DeleteInstallment(context.Context, int) error
}

type loanService struct {
	loanRepo repository.LoanRepository
}

func NewLoanService(db *gorm.DB) LoanService {
	return &loanService{
		loanRepo: repository.NewLoanRepository(db),
	}
}

func (s *loanService) Create(ctx context.Context, req dto.CreateLoanRequest) error {
	loanType, err := s.loanRepo.GetLoanType(ctx, req.LoanTypeID)
	if err != nil {
		return dto.ErrorWrap("Create loan: failed get loan type", err)
	}
	loan := req.WithInterest(loanType).ToLoanEntity()
	tokenInfo := ctx.Value(middleware.TokenInfoContextKey).(token.Claims)

	loan.CreatedBy = &tokenInfo.UserID
	err = s.loanRepo.Create(ctx, loan)
	if err != nil {
		return dto.ErrorWrap("failed create loan", err)
	}
	return nil
}

func (s *loanService) List(ctx context.Context, req dto.GetLoanListRequest) (dto.GetLoanListResponse, error) {
	var result dto.GetLoanListResponse
	var loans []dto.GetLoanResponse
	res, err := s.loanRepo.List(ctx, req)
	if err != nil {
		return result, dto.ErrorWrap("failed get Loan list", err)
	}

	for _, v := range res.Loans {
		tmp := dto.GetLoanResponse{}
		tmp.FromEntity(v)
		loans = append(loans, tmp)
	}

	result.Loans = loans
	result.Pagination = res.Pagination
	return result, nil
}

func (s *loanService) Get(ctx context.Context, ID int) (entity.Loan, error) {
	res, err := s.loanRepo.Get(ctx, ID)
	if err != nil {
		return res, dto.ErrorWrap("failed get Loan", err)
	}
	return res, nil
}
func (s *loanService) Update(ctx context.Context, req dto.UpdateLoanRequest) error {
	err := s.loanRepo.Update(ctx, req)
	if err != nil {
		return dto.ErrorWrap("failed update Loan", err)
	}
	return nil
}
func (s *loanService) Delete(ctx context.Context, ID int) error {
	err := s.loanRepo.Delete(ctx, ID)
	if err != nil {
		return dto.ErrorWrap("failed update Loan", err)
	}
	return nil
}
func (s *loanService) CreateLoanType(ctx context.Context, req dto.CreateLoanTypeRequest) error {
	loanType := entity.LoanType{
		Name:            req.Name,
		RatioPercentage: req.RatioPercentage,
	}
	tokenInfo := ctx.Value(middleware.TokenInfoContextKey).(token.Claims)
	loanType.CreatedBy = &tokenInfo.UserID

	err := s.loanRepo.CreateLoanType(ctx, loanType)
	if err != nil {
		return dto.ErrorWrap("failed create loan type", err)
	}
	return nil
}
func (s *loanService) UpdateLoanType(ctx context.Context, req dto.UpdateLoanTypeRequest) error {
	err := s.loanRepo.UpdateLoanType(ctx, req)
	if err != nil {
		return dto.ErrorWrap("failed update loan type", err)
	}
	return nil
}
func (s *loanService) ListLoanType(ctx context.Context) ([]dto.GetLoanTypeResponse, error) {
	rsp := []dto.GetLoanTypeResponse{}
	res, err := s.loanRepo.ListLoanType(ctx)
	if err != nil {
		return rsp, dto.ErrorWrap("failed get list loan type", err)
	}

	for _, v := range res {
		tmp := dto.GetLoanTypeResponse{}
		tmp.FromEntity(v)
		rsp = append(rsp, tmp)
	}

	return rsp, nil
}
func (s *loanService) CreateLoanInstallment(ctx context.Context, req dto.CreateLoanInstallmentRequest) error {
	loanInstallment := req.ToLoanInstallmentEntity()
	tokenInfo := ctx.Value(middleware.TokenInfoContextKey).(token.Claims)

	loanInstallment.CreatedBy = &tokenInfo.UserID
	err := s.loanRepo.CreateLoanInstallment(ctx, loanInstallment)
	if err != nil {
		return dto.ErrorWrap("failed create loan installment", err)
	}
	return nil
}
func (s *loanService) GetLoanInstallment(ctx context.Context, ID int) (entity.LoanInstallment, error) {
	res, err := s.loanRepo.GetLoanInstallment(ctx, ID)
	if err != nil {
		return res, dto.ErrorWrap("failed get loan installment", err)
	}
	return res, nil
}
func (s *loanService) ListLoanInstallment(ctx context.Context, req dto.GetLoanInstallmentListRequest) (dto.GetLoanInstallmentListResponse, error) {
	var result dto.GetLoanInstallmentListResponse

	res, err := s.loanRepo.ListLoanInstallment(ctx, req)
	if err != nil {
		return result, dto.ErrorWrap("failed get loan installment list", err)
	}

	for _, v := range res.LoanInstallments {
		tmp := dto.GetLoanInstallmentResponse{}
		tmp.FromEntity(v)
		result.LoanInstallments = append(result.LoanInstallments, tmp)
	}
	result.Pagination = res.Pagination
	return result, nil
}

func (s *loanService) UpdateLoanInstallment(ctx context.Context, req dto.UpdateLoanInstallmentRequest) error {
	err := s.loanRepo.UpdateLoanInstallment(ctx, req)
	if err != nil {
		return dto.ErrorWrap("failed update loan installment", err)
	}
	return nil
}
func (s *loanService) DeleteInstallment(ctx context.Context, ID int) error {
	err := s.loanRepo.DeleteLoanInstallment(ctx, ID)
	if err != nil {
		return dto.ErrorWrap("failed delete loan installment", err)
	}
	return nil
}
