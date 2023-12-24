package saving

import (
	"context"

	"github.com/jakskal/koperasi-v2/internal/entity"
	savingRepo "github.com/jakskal/koperasi-v2/internal/service/saving/repository"
	"github.com/jakskal/koperasi-v2/pkg/dto"
	"github.com/jakskal/koperasi-v2/pkg/middleware"
	"github.com/jakskal/koperasi-v2/pkg/token"
	"gorm.io/gorm"
)

type SavingService interface {
	Create(context.Context, dto.CreateSavingRequest) error
	List(context.Context, dto.GetSavingListRequest) (dto.GetSavingListResponse, error)
	Get(ctx context.Context, ID int) (entity.Saving, error)
	Update(context.Context, dto.UpdateSavingRequest) error
	Delete(ctx context.Context, ID int) error
	CreateSavingType(ctx context.Context, req dto.CreateSavingTypeRequest) error
	UpdateSavingType(context.Context, dto.UpdateSavingTypeRequest) error
	ListSavingType(context.Context) ([]entity.SavingType, error)
}

type savingService struct {
	savingRepo savingRepo.SavingRepository
}

func NewSavingService(db *gorm.DB) SavingService {
	return &savingService{
		savingRepo: savingRepo.NewSavingRepository(db),
	}
}

func (s *savingService) Create(ctx context.Context, req dto.CreateSavingRequest) error {
	saving := req.ToSavingEntity()
	tokenInfo := ctx.Value(middleware.TokenInfoContextKey).(token.Claims)
	saving.CreatedBy = &tokenInfo.UserID
	err := s.savingRepo.Create(ctx, saving)
	if err != nil {
		return dto.ErrorWrap("failed create saving", err)
	}
	return nil
}

func (s *savingService) List(ctx context.Context, req dto.GetSavingListRequest) (dto.GetSavingListResponse, error) {
	res, err := s.savingRepo.List(ctx, req)
	if err != nil {
		return dto.GetSavingListResponse{}, dto.ErrorWrap("failed get saving list", err)
	}
	return res, nil
}

func (s *savingService) Get(ctx context.Context, id int) (entity.Saving, error) {
	res, err := s.savingRepo.Get(ctx, id)
	if err != nil {
		return entity.Saving{}, dto.ErrorWrap("failed get saving", err)
	}
	return res, nil
}

func (s *savingService) Update(ctx context.Context, req dto.UpdateSavingRequest) error {
	err := s.savingRepo.Update(ctx, req)
	if err != nil {
		return dto.ErrorWrap("failed update saving", err)
	}
	return nil
}

func (s *savingService) Delete(ctx context.Context, ID int) error {
	err := s.savingRepo.Delete(ctx, ID)
	if err != nil {
		return dto.ErrorWrap("failed delete saving", err)
	}
	return nil
}

func (s *savingService) CreateSavingType(ctx context.Context, req dto.CreateSavingTypeRequest) error {
	tokenInfo := ctx.Value(middleware.TokenInfoContextKey).(token.Claims)

	err := s.savingRepo.CreateSavingTypes(ctx, entity.SavingType{
		Name: req.Name,
		TimeDefault: entity.TimeDefault{
			CreatedBy: &tokenInfo.UserID,
		},
	})
	if err != nil {
		return dto.ErrorWrap("failed create saving type", err)
	}

	return nil
}

func (s *savingService) UpdateSavingType(ctx context.Context, req dto.UpdateSavingTypeRequest) error {

	err := s.savingRepo.UpdateSavingType(ctx, req)
	if err != nil {
		return dto.ErrorWrap("failed update saving type", err)
	}
	return nil
}

func (s *savingService) ListSavingType(ctx context.Context) ([]entity.SavingType, error) {
	res, err := s.savingRepo.ListSavingType(ctx)
	if err != nil {
		return nil, dto.ErrorWrap("failed get saving type", err)
	}
	return res, nil
}
