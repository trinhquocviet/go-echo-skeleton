package model

import (
	"time"
)

// LoanDetail the data transfer object from database
type LoanDetail struct {
	ID         int64     `gorm:"column:id"`
	LoanID     int64     `gorm:"column:loan_id"`
	BookCopyID int64     `gorm:"column:book_copy_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}
