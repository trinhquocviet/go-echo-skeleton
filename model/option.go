package model

// Option the data transfer object from database
type Option struct {
  ID    int64  `gorm:"column:id"`
  Key   string `gorm:"column:key"`
  Value string `gorm:"column:value"`
}
