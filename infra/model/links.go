package model

type Link struct {
	ID       uint   `gorm:"primarykey"`
	Code     string `gorm:"uniqueIndex"`
	Link     string
	CreateAt int64
}
