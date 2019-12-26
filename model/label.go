package model

// Label the data transfer object from database
type Label struct {
	ID       int64  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Slug     string `gorm:"column:slug"`
	ParentID int64  `gorm:"column:parent_id"`
}
