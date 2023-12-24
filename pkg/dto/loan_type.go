package dto

import (
	"github.com/jakskal/koperasi-v2/internal/entity"
	"github.com/shopspring/decimal"
)

type CreateLoanTypeRequest struct {
	Name            string          `json:"name" binding:"required"`
	RatioPercentage decimal.Decimal `json:"ratio_percentage"`
}

type UpdateLoanTypeRequest struct {
	ID              int
	Name            string          `json:"name"`
	RatioPercentage decimal.Decimal `json:"ratio_percentage"`
}

type GetLoanTypeResponse struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	RatioPercentage    float64 `json:"ratio_percentage"`
	entity.TimeDefault `gorm:"embedded"`
}

func (l *GetLoanTypeResponse) FromEntity(req entity.LoanType) {
	l.ID = req.ID
	l.Name = req.Name
	l.RatioPercentage = req.RatioPercentage.InexactFloat64()
	l.TimeDefault = req.TimeDefault
}
