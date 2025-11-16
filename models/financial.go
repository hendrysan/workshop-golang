package models

import (
	"time"

	"gorm.io/gorm"
)

type FinancialTable struct {
}

func (FinancialTable) TableName() string {
	return "financials"
}

type Financial struct {
	FinancialTable 
	Id          int            `json:"id"`
	Category    string         `json:"category"`
	Nominal     int            `json:"nominal"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type FinancialCreateRequest struct {
	Category    string `json:"category"`
	Nominal     int    `json:"nominal"`
	Description string `json:"description"`
}
