package entity

import "time"

type SavingTransactionType int

const (
	SAVING_TRX_TYPE_DEPOSIT SavingTransactionType = iota + 1
	SAVING_TRX_TYPE_WITHDRAWAL
)

type SavingType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	TimeDefault `gorm:"embedded"`
}

type Saving struct {
	ID                int
	UserID            int
	SavingTypeID      int
	TransactionTypeID SavingTransactionType
	Amount            int
	Notes             string
	TimeDefault       `gorm:"embedded"`
	SavingChanges     []SavingChange `gorm:"foreignKey:SavingID;references:ID"`
}

type SavingChange struct {
	ID                int
	SavingID          int
	TransactionTypeID SavingTransactionType
	Notes             string
	Amount            int
	ChangesNotes      string
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         int       `json:"created_by"`
}
