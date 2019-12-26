package model

import (
  "time"
)

// Loan the data transfer object from database
type Loan struct {
  ID            int64       `gorm:"column:id"`
  Code          string      `gorm:"column:code"`
  MemberID      int64       `gorm:"column:member_id"`
  // BookCopyID    int64       `gorm:"column:book_copy_id"`
  Status        int         `gorm:"column:status"`
  CreatedAt     time.Time   `gorm:"column:created_at"`
  UpdatedAt     *time.Time  `gorm:"column:updated_at"`
}