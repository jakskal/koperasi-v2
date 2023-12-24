package entity

import (
	"time"
)

type TimeDefault struct {
	CreatedAt time.Time  `json:"created_at" gorm:"<-:create"`
	CreatedBy *int       `json:"created_by" gorm:"<-:create"`
	UpdatedAt time.Time  `json:"updated_at"`
	UpdatedBy *int       `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *int       `json:"deleted_by"`
}
