package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type LoanTransactionType int

const (
	LOAN_TRX_TYPE_ADMIN SavingTransactionType = iota + 1
	LOAN_TRX_TYPE_INSTALLMENT
)

type LoanType struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	RatioPercentage decimal.Decimal `json:"ratio_percentage"`
	TimeDefault     `gorm:"embedded"`
}

type Loan struct {
	ID                   int
	UserID               int
	LoanTypeID           int
	Name                 string
	Amount               decimal.Decimal
	RatioPercentage      decimal.Decimal
	TotalRatioAmount     decimal.Decimal
	InstallmentQtyTarget int
	Notes                string
	TimeDefault          `gorm:"embedded"`
	LoanChanges          []LoanChange `gorm:"foreignKey:LoanID;references:ID"`
}

type LoanChange struct {
	ID               int             `json:"id"`
	LoanID           int             `json:"loan_id"`
	Name             string          `json:"name"`
	Amount           decimal.Decimal `json:"amount"`
	RatioPercentage  decimal.Decimal `json:"ratio_percentage"`
	TotalRatioAmount decimal.Decimal `json:"total_ratio_amount"`
	Notes            string          `json:"notes"`
	ChangesNotes     string          `json:"change_notes"`
	CreatedAt        time.Time       `json:"created_at"`
	CreatedBy        int             `json:"created_by"`
}

type LoanInstallment struct {
	ID                     int
	LoanID                 int
	TransactionTypeID      LoanTransactionType
	PaymentDate            time.Time
	PrincipalAmount        decimal.Decimal
	InterestAmount         decimal.Decimal
	Notes                  string
	TimeDefault            `gorm:"embedded"`
	LoanInstallmentChanges []LoanInstallmentChange `gorm:"foreignKey:LoanInstallmentID;references:ID"`
}

type LoanInstallmentChange struct {
	ID                int
	LoanInstallmentID int
	TransactionTypeID LoanTransactionType
	PaymentDate       time.Time
	PrincipalAmount   decimal.Decimal
	InterestAmount    decimal.Decimal
	Notes             string
	ChangesNotes      string
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         int       `json:"created_by"`
}
