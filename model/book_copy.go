package model

import (
  "time"
)

// BookCopy the data transfer object from database
type BookCopy struct {
  ID            int64       `gorm:"column:id"`
  Code          string      `gorm:"column:code"`
  BookID        int64       `gorm:"column:book_id"`
  Status        int         `gorm:"column:status"`
  CreatedAt     time.Time   `gorm:"column:created_at"`
  UpdatedAt     *time.Time  `gorm:"column:updated_at"`
  DeletedAt     *time.Time  `gorm:"column:deleted_at"`
}

// TableName define name for gorm Table
func (inst *BookCopy) TableName() string {
	return "book_copies"
}