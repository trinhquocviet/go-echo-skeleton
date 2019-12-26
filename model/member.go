package model

import (
  "time"
)

// Member the data transfer object from database
type Member struct {
  ID          int64      `gorm:"column:id"`
  UID         string     `gorm:"column:uid"`
  DisplayName string     `gorm:"column:display_name"`
  CreatedAt   time.Time  `gorm:"column:created_at"`
  UpdatedAt   *time.Time `gorm:"column:updated_at"`
  DeletedAt   *time.Time `gorm:"column:deleted_at"`
}
