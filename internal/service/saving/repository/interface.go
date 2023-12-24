package repository

import (
	"context"

	"github.com/jakskal/koperasi-v2/internal/entity"
	"github.com/jakskal/koperasi-v2/pkg/dto"
)

type SavingRepository interface {
	Create(context.Context, entity.Saving) error
	Get(ctx context.Context, ID int) (entity.Saving, error)
	GetUserSavingDetail(ctx context.Context, ID int, userID int) (entity.Saving, error)
	List(ctx context.Context, req dto.GetSavingListRequest) (dto.GetSavingListResponse, error)
	Update(context.Context, dto.UpdateSavingRequest) error
	Delete(ctx context.Context, ID int) error
	CreateSavingTypes(ctx context.Context, req entity.SavingType) error
	UpdateSavingType(ctx context.Context, req dto.UpdateSavingTypeRequest) error
	ListSavingType(ctx context.Context) ([]entity.SavingType, error)
}
