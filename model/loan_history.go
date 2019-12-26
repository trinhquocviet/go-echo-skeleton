package model

import (
  "time"
)

// LoanHistory the data transfer object from database
type LoanHistory struct {
  ID            int64       `gorm:"column:id"`
  LoanID        int64       `gorm:"column:loan_id"`
  Status        int         `gorm:"column:status"`
  CreatedAt     time.Time   `gorm:"column:created_at"`
}

// TableName define name for gorm Table
func (inst *LoanHistory) TableName() string {
	return "loan_histories"
}