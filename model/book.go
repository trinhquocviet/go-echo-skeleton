package model

import (
	"time"
)

// Book the data transfer object from database
type Book struct {
	ID            int64      `gorm:"column:id"`
	ISBN          string     `gorm:"column:isbn"`
	Title         string     `gorm:"column:title"`
	DatePublished time.Time  `gorm:"column:date_published"`
	CreatedAt     time.Time  `gorm:"column:created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at"`
}
