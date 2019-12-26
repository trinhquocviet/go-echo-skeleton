package model

// BookExtend the data transfer object from database
type BookExtend struct {
  ID            int64       `gorm:"column:id"`
  BookID        int64       `gorm:"column:book_id"`
  Key           string      `gorm:"column:key"`
  Value         string      `gorm:"column:value"`
}
