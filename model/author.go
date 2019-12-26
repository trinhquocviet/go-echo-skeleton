package model

import (
  "time"
)

// Author the data transfer object from database
type Author struct {
  ID          int64       `gorm:"column:id"`
  Name        string      `gorm:"column:name"`
  CreatedAt   time.Time   `gorm:"column:created_at"`
  UpdatedAt   *time.Time  `gorm:"column:updated_at"`
}
