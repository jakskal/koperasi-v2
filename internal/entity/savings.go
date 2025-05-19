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
	ID                int                   `json:"id"`
	UserID            int                   `json:"user_id"`
	SavingTypeID      int                   `json:"saving_type_id"`
	TransactionTypeID SavingTransactionType `json:"transaction_type_id"`
	TransactionDate   time.Time             `json:"transaction_date"`
	Amount            int                   `json:"amount"`
	Notes             string                `json:"notes"`
	TimeDefault       `gorm:"embedded"`
	SavingChanges     []SavingChange `gorm:"foreignKey:SavingID;references:ID"`
	User              User           `gorm:"foreignKey:ID;references:UserID" json:"user"`
	SavingType        SavingType     `gorm:"foreignKey:ID;references:SavingTypeID" json:"saving_type"`
}

type SavingChange struct {
	ID                int
	SavingID          int
	TransactionTypeID SavingTransactionType
	TransactionDate   time.Time
	Notes             string
	Amount            int
	ChangesNotes      string
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         int       `json:"created_by"`
}
