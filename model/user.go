package model

import (
  "time"
)

// User the data transfer object from database
type User struct {
  ID          string     `gorm:"column:id"`
  Email       string     `gorm:"column:email"`
  Password    string     `gorm:"column:password"`
  Token       string     `gorm:"column:token"`
  CreatedAt   time.Time  `gorm:"column:created_at"`
  UpdatedAt   *time.Time `gorm:"column:updated_at"`
  DeletedAt   *time.Time `gorm:"column:deleted_at"`
}
